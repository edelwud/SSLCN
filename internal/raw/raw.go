package raw

import (
	"golang.org/x/net/ipv4"
	"net"
)

type Raw struct {
	*ipv4.PacketConn
	address string
}

func New(address string) (*Raw, error) {
	socket, err := net.ListenPacket("ip4:1", address)
	if err != nil {
		return nil, err
	}

	p := ipv4.NewPacketConn(socket)

	return &Raw{p, address}, nil
}
