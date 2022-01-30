package ping

import (
	"SSLCN/internal/protocol/icmp"
	"bytes"
	"golang.org/x/net/ipv4"
	"io"
	"log"
	"net"
)

type Ping struct {
	address *net.IPAddr
	socket  *ipv4.PacketConn
}

func (p Ping) Write(buf []byte) (n int, err error) {
	n, err = p.socket.WriteTo(buf, nil, p.address)
	if err != nil {
		return 0, err
	}

	return len(buf), nil
}

func (p Ping) Read(buf []byte) (n int, err error) {
	n, _, _, err = p.socket.ReadFrom(buf)
	if err != nil {
		return 0, err
	}

	return n, io.EOF
}

func (p Ping) Send() error {
	payload := new(bytes.Buffer)
	payload.Write([]byte("HEY"))

	packed, err := icmp.Pack(icmp.EchoReply, payload)
	if err != nil {
		return err
	}

	_, err = io.Copy(p, packed)
	if err != nil {
		return err
	}

	return nil
}

func (p Ping) Receive() (*bytes.Buffer, error) {
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

	return &Ping{dst, p}, nil
}
