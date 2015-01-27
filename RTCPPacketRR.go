package rtp

type RTCPPacketRR interface {
	RTCPPacket

	GetSenderSSRC() uint32
	SetSenderSSRC(uint32)

	GetReportBlock(n byte) RTCPReportBlock
	SetReportBlock(n byte, rr RTCPReportBlock)

	GetExtension() []byte
	SetExtension([]byte)
}
