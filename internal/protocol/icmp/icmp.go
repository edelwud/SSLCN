package icmp

import (
	"SSLCN/pkg/checksum"
	"bytes"
	"encoding/binary"
	"os"
)

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	PacketID    uint16
	SequenceNum uint16
}

var (
	EchoReply = ICMP{Type: 0, Code: 0}
)

func Pack(header ICMP, payload *bytes.Buffer) (*bytes.Buffer, error) {
	message := new(bytes.Buffer)

	header.PacketID = uint16(os.Getpid()) & 0xFFFF

	err := binary.Write(message, binary.BigEndian, header)
	if err != nil {
		return nil, err
	}

	header.Checksum = checksum.CalculateInternetChecksum(message)

	message.Reset()

	err = binary.Write(message, binary.BigEndian, header)
	if err != nil {
		return nil, err
	}

	_, err = payload.WriteTo(message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func Unpack(message *bytes.Buffer) (header ICMP, payload *bytes.Buffer, err error) {
	err = binary.Read(message, binary.BigEndian, &header)
	if err != nil {
		return ICMP{}, nil, err
	}

	return header, message, err
}
