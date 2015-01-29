package rtp

type RTCPSDESType byte

const (
	RTCP_SDES_END RTCPSDESType = iota
	RTCP_SDES_CNAME
	RTCP_SDES_NAME
	RTCP_SDES_EMAIL
	RTCP_SDES_PHONE
	RTCP_SDES_LOC
	RTCP_SDES_TOOL
	RTCP_SDES_NOTE
	RTCP_SDES_PRIV
)

type RTCPPacketSDES interface {
	RTCPPacket

	GetSDESItem(n byte) RTCPSDESItem
	SetSDESItem(n byte, sdes RTCPSDESItem)
}

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
