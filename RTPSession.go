package rtp

import (
	"log"
	"net"
	"sync"
	"time"
)

type RTPSession interface {
	AddRTPListener(RTPListener)
	RemoveRTPListener(RTPListener)

	AddRTCPListener(RTCPListener)
	RemoveRTCPListener(RTCPListener)

	SetRTPTransport(Transport)
	SetRTCPTransport(Transport)

	Run() error
	Stop()

	SetCName(string)
	SetEmail(string)
	SetPhone(string)
	SetLocation(string)
	SetTool(string)
	SetNote(string)

	SendRTPPacket(RTPPacket)
	SendRTCPPacket(RTCPPacket)
}

type TransportType byte

const (
	TRANSPORT_RTP TransportType = iota
	TRANSPORT_RTCP
	TRANSPORT_NUM
)

type session struct {
	rtpListeners  map[RTPListener]RTPListener
	rtcpListeners map[RTCPListener]RTCPListener
	transports    [TRANSPORT_NUM]Transport

	quit      chan bool
	waitGroup *sync.WaitGroup
}

func NewSession() *session {
	this := &session{}

	this.rtpListeners = make(map[RTPListener]RTPListener)
	this.rtcpListeners = make(map[RTCPListener]RTCPListener)

	this.quit = make(chan bool)
	this.waitGroup = &sync.WaitGroup{}

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

func (this *session) SetRTPTransport(t Transport) {
	this.transports[TRANSPORT_RTP] = t
}

func (this *session) SetRTCPTransport(t Transport) {
	this.transports[TRANSPORT_RTCP] = t
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
