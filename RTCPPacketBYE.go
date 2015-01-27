package rtp

type RTCPPacketBYE interface {
	RTCPPacket

	GetSCSRC() []uint32
	SetSCSRC([]uint32)

	GetReasonLength() byte
	SetReasonLength(byte)

	GetReason() []byte
	SetReason([]byte)
}
