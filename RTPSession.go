package rtp

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

type Session interface {
	AddRTPListener(RTPListener)
	RemoveRTPListener(RTPListener)

	AddRTCPListener(RTCPListener)
	RemoveRTCPListener(RTCPListener)

	SendRTPPacket(RTPPacket, addr *net.UDPAddr) (int, error)
	SendRTCPPacket(RTCPPacket, addr *net.UDPAddr) (int, error)

	Run() error
	Stop()
}

type TransportType byte

const (
	TRANSPORT_RTP TransportType = iota
	TRANSPORT_RTCP
	TRANSPORT_NUM
)

type SessionParam struct {
	cname    string
	email    string
	phone    string
	location string
	tool     string
	note     string
}

type session struct {
	rtpListeners  map[RTPListener]RTPListener
	rtcpListeners map[RTCPListener]RTCPListener
	transports    [TRANSPORT_NUM]Transport

	quit      chan bool
	waitGroup *sync.WaitGroup

	param SessionParam
}

func NewSession(rtpt, rtcpt Transport, sp SessionParam) *session {
	this := &session{}

	this.rtpListeners = make(map[RTPListener]RTPListener)
	this.rtcpListeners = make(map[RTCPListener]RTCPListener)

	this.quit = make(chan bool)
	this.waitGroup = &sync.WaitGroup{}

	this.transports[TRANSPORT_RTP] = rtpt
	this.transports[TRANSPORT_RTCP] = rtcpt
	this.param = sp

	return this
}

func (this *session) AddRTPListener(l RTPListener) {
	this.rtpListeners[l] = l
}

func (this *session) RemoveRTPListener(l RTPListener) {
	delete(this.rtpListeners, l)
}

func (this *session) AddRTCPListener(l RTCPListener) {
	this.rtcpListeners[l] = l
}

func (this *session) RemoveRTCPListener(l RTCPListener) {
	delete(this.rtcpListeners, l)
}

func (this *session) SendRTPPacket(pkt RTPPacket, addr *net.UDPAddr) (int, error) {
	if this.transports[TRANSPORT_RTP].GetConn() != nil {
		return this.transports[TRANSPORT_RTP].GetConn().WriteToUDP(pkt.Bytes(), addr)
	}

	return 0, errors.New("Please Run Session first!")
}

func (this *session) SendRTCPPacket(pkt RTCPPacket, addr *net.UDPAddr) (int, error) {
	if this.transports[TRANSPORT_RTCP].GetConn() != nil {
		return this.transports[TRANSPORT_RTCP].GetConn().WriteToUDP(pkt.Bytes(), addr)
	}

	return 0, errors.New("Please Run Session first!")
}

func (this *session) Run() error {
	for tt, tp := range this.transports {
		var conn *net.UDPConn
		var err error

		if tp.IsMulticast() {
			conn, err = net.ListenMulticastUDP(tp.GetNetwork(), tp.GetInterface(), tp.GetUDPAddr())
		} else {
			conn, err = net.ListenUDP(tp.GetNetwork(), tp.GetUDPAddr())
		}
		if err != nil {
			close(this.quit)
			return err
		}
		this.waitGroup.Add(1)
		tp.SetConn(conn)
		go this.ServeConn(conn, TransportType(tt))
	}

	return nil
}

func (this *session) Stop() {
	close(this.quit)
	this.waitGroup.Wait()
}

//////////////////////////////////////////////////////////////////////

func (this *session) ServeConn(conn *net.UDPConn, tt TransportType) {
	defer this.waitGroup.Done()
	defer conn.Close()

	var n int
	var addr *net.UDPAddr
	var err error
	var buf [RTP_MTU_SIZE]byte

	for {
		select {
		case <-this.quit:
			log.Println("Disconnecting", conn.RemoteAddr())
			return
		default:
		}

		conn.SetReadDeadline(time.Now().Add(1e9)) //wait for 1 second
		if n, addr, err = conn.ReadFromUDP(buf[:]); err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}
		} else {
			if tt == TRANSPORT_RTP {
				if pkt, err := RTPPacketize(buf[0:n]); err != nil {
					continue
				} else {
					for _, ln := range this.rtpListeners {
						ln.HandleRTPEvent(pkt, addr)
					}
				}
			} else {
				if pkt, err := RTCPPacketize(buf[0:n]); err != nil {
					continue
				} else {
					for _, ln := range this.rtcpListeners {
						ln.HandleRTCPEvent(pkt, addr)
					}
				}
			}
		}

	}
}
