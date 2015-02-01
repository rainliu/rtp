package rtp

import (
	"bytes"
)

const (
	RTP_VERSION  = 2
	RTP_SEQ_MOD  = 1 << 16
	RTP_MAX_SDES = 255
	RTP_MTU_SIZE = 1500
)

type RTPPacket interface {
	Bytes() []byte
	Parse([]byte) error

	GetVersion() byte
	SetVersion(byte)

	GetPadding() bool
	SetPadding(bool)

	GetExtension() bool
	SetExtension(bool)

	GetCSRCCount() byte
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
	Bytes() []byte

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

type rtp_packet struct {
	version        byte
	padding        bool
	extension      bool
	count          byte
	marker         bool
	payloadType    byte
	sequenceNumber uint16
	timeStamp      uint32
	ssrc           uint32
	csrcs          []uint32

	headerExtension RTPHeaderExtension
	payload         []byte
}

func NewRTPPacket() *rtp_packet {
	this := &rtp_packet{}
	return this
}

func (this *rtp_packet) Bytes() []byte {
	var buffer bytes.Buffer

	//buffer.WriteByte((byte(this.packetType) << 4) | (this.packetFlag & 0x0F))
	//buffer.WriteByte(0)

	return buffer.Bytes()
}

func (this *rtp_packet) Parse(buffer []byte) error {
	// if buffer == nil || len(buffer) != 2 {
	// 	return fmt.Errorf("Invalid %s Control Packet Size %x\n", PACKET_TYPE_STRINGS[this.packetType], len(buffer))
	// }

	// if packetType := PacketType((buffer[0] >> 4) & 0x0F); packetType != this.packetType {
	// 	return fmt.Errorf("Invalid %s Control Packet Type %x\n", PACKET_TYPE_STRINGS[this.packetType], packetType)
	// }
	// if packetFlag := buffer[0] & 0x0F; packetFlag != this.packetFlag {
	// 	return fmt.Errorf("Invalid %s Control Packet Flags %x\n", this.packetType, packetFlag)
	// }
	// if buffer[1] != 0 {
	// 	return fmt.Errorf("Invalid %s Control Packet Remaining Length %x\n", PACKET_TYPE_STRINGS[this.packetType], buffer[1])
	// }

	return nil
}

func (this *rtp_packet) GetVersion() byte {
	return this.version
}
func (this *rtp_packet) SetVersion(version byte) {
	this.version = version
}

func (this *rtp_packet) GetPadding() bool {
	return this.padding
}
func (this *rtp_packet) SetPadding(padding bool) {
	this.padding = padding
}

func (this *rtp_packet) GetExtension() bool {
	return this.extension
}
func (this *rtp_packet) SetExtension(extension bool) {
	this.extension = extension
}

func (this *rtp_packet) GetCSRCCount() byte {
	return this.count
}
func (this *rtp_packet) SetCSRCCount(count byte) {
	this.count = count
}

func (this *rtp_packet) GetMarker() bool {
	return this.marker
}
func (this *rtp_packet) SetMarker(marker bool) {
	this.marker = marker
}

func (this *rtp_packet) GetPayloadType() byte {
	return this.payloadType
}
func (this *rtp_packet) SetPayloadType(payloadType byte) {
	this.payloadType = payloadType
}

func (this *rtp_packet) GetSequenceNumber() uint16 {
	return this.sequenceNumber
}
func (this *rtp_packet) SetSequenceNumber(sequenceNumber uint16) {
	this.sequenceNumber = sequenceNumber
}

func (this *rtp_packet) GetTimeStamp() uint32 {
	return this.timeStamp
}
func (this *rtp_packet) SetTimeStamp(timeStamp uint32) {
	this.timeStamp = timeStamp
}

func (this *rtp_packet) GetSSRC() uint32 {
	return this.ssrc
}
func (this *rtp_packet) SetSSRC(ssrc uint32) {
	this.ssrc = ssrc
}

func (this *rtp_packet) GetCSRC() []uint32 {
	return this.csrcs
}
func (this *rtp_packet) SetCSRC(csrcs []uint32) {
	this.csrcs = csrcs
}

func (this *rtp_packet) GetHeaderExtension() RTPHeaderExtension {
	return this.headerExtension
}
func (this *rtp_packet) SetHeaderExtension(headerExtension RTPHeaderExtension) {
	this.headerExtension = headerExtension
}

func (this *rtp_packet) GetPayload() []byte {
	return this.payload
}
func (this *rtp_packet) SetPayload(payload []byte) {
	this.payload = payload
}
