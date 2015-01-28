package rtp

type RTPListener interface {
	HandleRTPEvent(RTPPacket)
}

type RTCPListener interface {
	HandleRTCPEvent(RTCPPacket)
}
