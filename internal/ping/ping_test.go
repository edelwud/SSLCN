package ping

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	p, err := New("google.com")
	if err != nil {
		panic(err)
	}
	defer p.Close()
}

func TestPing_Send(t *testing.T) {
	p, err := New("1.1.1.1")
	if err != nil {
		panic(err)
	}
	defer p.Close()

	err = p.Send()
	if err != nil {
		panic(err)
	}

	buf, err := p.Receive()
	if err != nil {
		panic(err)
	}

	spew.Dump(buf.Bytes())
}

func TestPing_Receive(t *testing.T) {
	p, err := New("google.com")
	if err != nil {
		panic(err)
	}
	defer p.Close()

	for i := 0; i < 100; i++ {
		go func() {
			err := p.Send()
			if err != nil {
				panic(err)
			}

			value, err := p.Receive()
			if err != nil {
				panic(err)
			}

			spew.Dump(value.Bytes())
		}()
	}

	time.Sleep(time.Second * 10)
}
