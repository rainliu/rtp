package rtp

import (
	"log"
	"net"
	"sync"
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

type session struct {
	rtpListeners  map[RTPListener]RTPListener
	rtcpListeners map[RTCPListener]RTCPListener
	transports    [2]Transport

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
	this.transports[0] = t
}

func (this *session) SetRTCPTransport(t Transport) {
	this.transports[1] = t
}

func (this *session) Run() error {
	for _, t := range this.transports {
		var conn *net.UDPConn
		var err error

		if t.IsMulticast() {
			conn, err = net.ListenMulticastUDP(t.GetNetwork(), t.GetInterface(), t.GetUDPAddr())
		} else {
			conn, err = net.ListenUDP(t.GetNetwork(), t.GetUDPAddr())
		}
		if err != nil {
			return err
		}
		this.waitGroup.Add(1)
		t.SetConn(conn)
		go this.ServeConn(conn)
	}

	return nil
}

func (this *session) Stop() {
	close(this.quit)
	this.waitGroup.Wait()
}

func (this *session) ServeConn(conn *net.UDPConn) {
	defer this.waitGroup.Done()
	defer conn.Close()

	//var buf []byte
	//var err error
	for {
		select {
		case <-this.quit:
			log.Println("Disconnecting", conn.RemoteAddr())
			return
		default:
		}

	}
}
