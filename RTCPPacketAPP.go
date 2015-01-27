package rtp

type RTCPPacketAPP interface {
	RTCPPacket

	GetSCSRC() uint32
	SetSCSRC(uint32)

	GetName() [4]byte
	SetName([4]byte)

	GetAppData() []byte
	SetAppData([]byte)
}
