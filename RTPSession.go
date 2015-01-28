package rtp

import (
	"fmt"
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

	for {
		select {
		case <-this.quit:
			log.Println("Disconnecting", conn.RemoteAddr())
			return
		default:
		}
		if buf, addr, err := this.ReadPacket(conn); err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}
		} else {
			if tt == TRANSPORT_RTP {
				if pkt, err := RTPPacketize(buf); err != nil {
					continue
				} else {
					for _, ln := range this.rtpListeners {
						ln.HandleRTPEvent(pkt, addr)
					}
				}
			} else {
				if pkt, err := RTCPPacketize(buf); err != nil {
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

func (this *session) ReadPacket(conn *net.UDPConn) ([]byte, *net.UDPAddr, error) {
	var pkt [4]byte
	var n int
	var addr1, addr2 *net.UDPAddr
	var err error

	conn.SetReadDeadline(time.Now().Add(1e9)) //wait for 1 second
	if n, addr1, err = conn.ReadFromUDP(pkt[:]); err != nil {
		return nil, addr1, err
	}
	if n != 4 {
		return nil, addr1, fmt.Errorf("Invalid RTP/RTCP Packet Fixed Header %d from %s", n, addr1.String())
	}

	length := (int(pkt[2]) << 8) | int(pkt[3])
	buf := make([]byte, length+4)
	copy(buf[0:4], pkt[:])

	if n, addr2, err = conn.ReadFromUDP(buf[4:]); err != nil {
		return nil, addr2, err
	}
	if n != length {
		return nil, addr2, fmt.Errorf("Invalid RTP/RTCP Packet Variable Length %d from %s", n, addr2.String())
	}
	if addr1.String() != addr2.String() {
		return nil, nil, fmt.Errorf("Invalid RTP/RTCP Packet Address between Fixed Header %s and Variable Part %s", addr1.String(), addr2.String())
	}

	return buf, addr2, nil
}
