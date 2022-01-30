package traceroute

import (
	"SSLCN/internal/protocol/icmp"
	"SSLCN/internal/raw"
	"bytes"
	"golang.org/x/net/ipv4"
	"net"
	"time"
)

type Traceroute struct {
	*raw.Raw
	address net.Addr
}

type Route struct {
	Index   int
	Address net.Addr
	Names   []string
	time.Duration
	*ipv4.ControlMessage
}

func (t Traceroute) Run(hops int) ([]Route, error) {
	routes := make([]Route, 0)

	for i := 1; i < hops; i++ {
		payload := new(bytes.Buffer)
		payload.Write([]byte("traceroute"))

		packet, err := icmp.Pack(icmp.NewEchoReply(0, uint16(i)), payload)
		if err != nil {
			return nil, err
		}

		err = t.SetTTL(i)
		if err != nil {
			return nil, err
		}

		begin := time.Now()

		_, err = t.WriteTo(packet.Bytes(), nil, t.address)
		if err != nil {
			return nil, err
		}

		err = t.SetReadDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			return nil, err
		}

		buf := make([]byte, 1024)
		_, cm, src, err := t.ReadFrom(buf)
		if err != nil {
			return nil, err
		}

		rtt := time.Since(begin)

		header, _, err := icmp.Unpack(bytes.NewBuffer(buf))
		if err != nil {
			return nil, err
		}

		switch header.Type {
		case uint8(ipv4.ICMPTypeTimeExceeded):
			names, _ := net.LookupAddr(src.String())
			routes = append(routes, Route{i, src, names, rtt, cm})
		case uint8(ipv4.ICMPTypeEchoReply):
			names, _ := net.LookupAddr(src.String())
			routes = append(routes, Route{i, src, names, rtt, cm})

			break
		}
	}

	return routes, nil
}

func New(socket *raw.Raw, destination string) (*Traceroute, error) {
	dst, err := net.ResolveIPAddr("ip4", destination)
	if err != nil {
		return nil, err
	}

	return &Traceroute{socket, dst}, nil
}
