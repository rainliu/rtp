package rtp

type RTCPPacketSR interface {
	RTCPPacket

	GetSenderSSRC() uint32
	SetSenderSSRC(uint32)

	GetNTPTimeStamp() uint64
	SetNTPTimeStamp(uint64)

	GetRTPTimeStamp() uint32
	SetRTPTimeStamp(uint32)

	GetSenderPacketCount() uint32
	SetSenderPacketCount(uint32)

	GetSenderOctetCount() uint32
	SetSenderOctetCount(uint32)

	GetReportBlock(n byte) RTCPReportBlock
	SetReportBlock(n byte, rr RTCPReportBlock)

	GetExtension() []byte
	SetExtension([]byte)
}
