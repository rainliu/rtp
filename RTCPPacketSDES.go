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

	GetType() RTCPSDESType
	SetType(RTCPSDESType)

	GetLength() byte
	SetLength(byte)

	GetValue() []byte
	SetValue([]byte)
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
