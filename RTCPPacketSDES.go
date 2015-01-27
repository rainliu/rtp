package rtp

type RTCPPacketSDES interface {
	RTCPPacket

	GetSDESItem(n byte) RTCPSDESItem
	SetSDESItem(n byte, sdes RTCPSDESItem)
}

type RTCPSDESType byte

const (
	CNAME RTCPSDESType = 1 + iota
	NAME
	EMAIL
	PHONE
	LOC
	TOOL
	NOTE
	PRIV
)

type RTCPSDESItem interface {
	GetSCSRC() uint32
	SetSCSRC(uint32)

	GetSDESType() RTCPSDESType
	SetSDESType(RTCPSDESType)

	GetLength() byte
	SetLength(byte)

	GetContent() []byte
	SetContent([]byte)
}

type RTCPSDESItemPriv interface {
	RTCPSDESItem

	GetPrefixLength() byte
	SetPrefixLength(byte)

	GetPrefixString() []byte
	SetPrefixString([]byte)

	GetValueString() []byte
	SetValueString([]byte)
}
