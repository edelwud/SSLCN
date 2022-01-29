package icmp

import (
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestPack(t *testing.T) {
	payload := new(bytes.Buffer)
	payload.Write([]byte("Test payload"))

	message, err := Pack(EchoReply, payload)
	if err != nil {
		panic(err)
	}

	spew.Dump(message.Bytes())
}

func TestUnpack(t *testing.T) {
	payload := new(bytes.Buffer)
	payload.Write([]byte("Test payload"))

	message, err := Pack(EchoReply, payload)
	if err != nil {
		panic(err)
	}

	spew.Dump(message.Bytes())

	header, payload2, err := Unpack(message)
	if err != nil {
		panic(err)
	}

	spew.Dump(header, payload2)
}
