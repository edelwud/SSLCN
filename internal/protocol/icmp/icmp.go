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
	EchoReply = ICMP{Type: 8, Code: 0, SequenceNum: 1}
)

func Pack(header ICMP, payload *bytes.Buffer) (*bytes.Buffer, error) {
	message := new(bytes.Buffer)

	header.PacketID = uint16(os.Getpid()) & 0xFFFF

	err := binary.Write(message, binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}

	message.Write(payload.Bytes())

	header.Checksum = checksum.CalculateInternetChecksum(message)

	message.Reset()

	err = binary.Write(message, binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}

	message.Write(payload.Bytes())

	return message, nil
}

func Unpack(message *bytes.Buffer) (header ICMP, payload *bytes.Buffer, err error) {
	err = binary.Read(message, binary.LittleEndian, &header)
	if err != nil {
		return ICMP{}, nil, err
	}

	return header, message, err
}
