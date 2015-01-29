package rtp

const (
	RTP_VERSION  = 2
	RTP_SEQ_MOD  = 1 << 16
	RTP_MAX_SDES = 255
	RTP_MTU_SIZE = 1500
)

type IBytizer interface {
	IBytize() []byte
}

type IParser interface {
	IParse([]byte) error
}

type RTPPacket interface {
	IBytizer
	Bytes() []byte

	IParser
	Parse([]byte) error

	GetVersion() byte
	SetVersion(byte)

	GetPadding() bool
	SetPadding(bool)

	GetExtension() bool
	SetExtension(bool)

	GetCSRCCount() []byte
	SetCSRCCount(byte)

	GetMarker() bool
	SetMarker(bool)

	GetPayloadType() byte
	SetPayloadType(byte)

	GetSequenceNumber() uint16
	SetSequenceNumber(uint16)

	GetTimeStamp() uint32
	SetTimeStamp(uint32)

	GetSSRC() uint32
	SetSSRC(uint32)

	GetCSRC() []uint32
	SetCSRC([]uint32)

	GetHeaderExtension() RTPHeaderExtension
	SetHeaderExtension(RTPHeaderExtension)

	GetPayload() []byte
	SetPayload([]byte)
}

type RTPHeaderExtension interface {
	IBytizer
	Bytes() []byte

	IParser
	Parse([]byte) error

	GetReserved() uint16
	SetReserved(uint16)

	GetLength() uint16
	SetLength(uint16)

	GetExtensionData() []byte
	SetExtensionData([]byte)
}

func RTPPacketize(buf []byte) (RTPPacket, error) {
	return nil, nil
}
