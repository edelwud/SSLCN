package icmp

import (
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

func TestPack(t *testing.T) {
	binary, err := time.Now().MarshalBinary()
	if err != nil {
		panic(err)
	}

	payload := new(bytes.Buffer)
	payload.Write(binary)

	message, err := Pack(EchoReply, payload)
	if err != nil {
		panic(err)
	}

	spew.Dump(message)
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
