package rtp

type RTCPPacketType byte

const (
	RTCP_SR RTCPPacketType = 200
	RTCP_RR
	RTCP_SDES
	RTCP_BYE
	RTCP_APP
)

const (
	RTCP_VALID_MASK  = (0xc000 | 0x2000 | 0x00fe)
	RTCP_VALID_VALUE = ((RTP_VERSION << 14) | int(RTCP_SR))
)

type IBytizer interface {
	IBytize() []byte
}

type IParser interface {
	IParse([]byte) error
}

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
