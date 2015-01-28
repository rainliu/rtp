package rtp

type RTCPPacketType byte

const (
	SR RTCPPacketType = 200
	RR
	SDES
	BYE
	APP
)

type RTCPPacket interface {
	IBytizer
	Bytes() []byte

	IParser
	Parse([]byte) error

	GetVersion() byte
	SetVersion(byte)

	GetPadding() bool
	SetPadding(bool)

	GetCount() byte
	SetCount(byte)

	GetPacketType() RTCPPacketType
	SetPacketType(RTCPPacketType)

	GetLength() uint16
	SetLength(uint16)
}

type RTCPReportBlock interface {
	GetSSRC() uint32
	SetSSRC(uint32)

	GetFractionLost() byte
	SetFractionLost(byte)

	GetCumulativeNumberOfPacketLost() uint32
	SetCumulativeNumberOfPacketLost(uint32)

	GetExtendedHighestSequenceNumberReceived() uint32
	SetExtendedHighestSequenceNumberReceived(uint32)

	GetInterarrivalJitter() uint32
	SetInterarrivalJitter(uint32)

	GetLSR() uint32
	SetLSR(uint32)

	GetDLSR() uint32
	SetDLSR(uint32)
}

func RTCPPacketize(buf []byte) (RTCPPacket, error) {
	return nil, nil
}
