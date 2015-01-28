package rtp

import "net"

type RTPListener interface {
	HandleRTPEvent(RTPPacket, *net.UDPAddr)
}

type RTCPListener interface {
	HandleRTCPEvent(RTCPPacket, *net.UDPAddr)
}
