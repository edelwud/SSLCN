package ping

import (
	"SSLCN/internal/protocol/icmp"
	"bytes"
	"golang.org/x/net/ipv4"
	"io"
	"log"
	"math/rand"
	"net"
	"syscall"
	"time"
)

type Ping struct {
	address      *net.IPAddr
	socket       *ipv4.PacketConn
	writtenBytes int
	identifier   uint16
}

func (p *Ping) Write(buf []byte) (n int, err error) {
	n, err = p.socket.PacketConn.WriteTo(buf, p.address)
	if err != nil {
		return 0, err
	}

	p.writtenBytes += n

	return n, nil
}

func (p *Ping) Read(buf []byte) (n int, err error) {
	oob := make([]byte, 2048)
	b := make([]byte, 2048)
	messages := make([]ipv4.Message, 0)
	messages = append(messages, ipv4.Message{
		Buffers: append(make([][]byte, 0), b),
		OOB:     oob,
		Addr:    p.address,
		N:       len(b),
		NN:      len(oob),
		Flags:   syscall.MSG_PEEK,
	})

	n, err = p.socket.ReadBatch(messages, syscall.MSG_PEEK)
	if err != nil {
		return 0, err
	}

	header, _, err := icmp.Unpack(bytes.NewBuffer(messages[0].Buffers[0][20:]))
	if err != nil {
		return 0, err
	}

	if header.PacketID == p.identifier {
		n, _, _, err = p.socket.ReadFrom(buf)
		if err != nil {
			return 0, err
		}

		p.writtenBytes -= n
		if p.writtenBytes == 0 {
			return n, io.EOF
		}
	}

	return n, io.EOF
}

func (p *Ping) Send() error {
	now := time.Now().String()

	payload := new(bytes.Buffer)
	payload.Write([]byte(now))

	packed, err := icmp.Pack(icmp.NewEchoReply(p.identifier, 1), payload)
	if err != nil {
		return err
	}

	_, err = io.Copy(p, packed)
	if err != nil {
		return err
	}

	return nil
}

func (p *Ping) Receive() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	_, err := io.Copy(buf, p)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (p Ping) Close() error {
	err := p.socket.Close()
	if err != nil {
		return err
	}

	return nil
}

func New(destination string) (*Ping, error) {
	dst, err := net.ResolveIPAddr("ip4", destination)
	if err != nil {
		log.Fatal(err)
	}

	socket, err := net.ListenPacket("ip4:1", "0.0.0.0")
	if err != nil {
		return nil, err
	}

	p := ipv4.NewPacketConn(socket)

	err = p.SetControlMessage(ipv4.FlagTTL|ipv4.FlagSrc|ipv4.FlagDst|ipv4.FlagInterface, true)
	if err != nil {
		return nil, err
	}

	return &Ping{dst, p, 0, uint16(rand.Intn(65536))}, nil
}
