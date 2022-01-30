package ping

import (
	"SSLCN/internal/raw"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	r, err := raw.New("0.0.0.0")
	if err != nil {
		panic(err)
	}

	p, err := New(r, "google.com")
	if err != nil {
		panic(err)
	}
	defer p.Close()
}

func TestPing_Send(t *testing.T) {
	r, err := raw.New("0.0.0.0")
	if err != nil {
		panic(err)
	}

	p, err := New(r, "1.1.1.1")
	if err != nil {
		panic(err)
	}
	defer p.Close()

	err = p.Send(1)
	if err != nil {
		panic(err)
	}

	_, buf, err := p.Receive()
	if err != nil {
		panic(err)
	}

	spew.Dump(buf.Bytes())
}

func TestPing_Run(t *testing.T) {
	r, err := raw.New("0.0.0.0")
	if err != nil {
		panic(err)
	}

	p, err := New(r, "1.1.1.1")
	if err != nil {
		panic(err)
	}
	defer p.Close()

	run, err := p.Run(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(run)
}

func TestPing_Receive(t *testing.T) {
	r, err := raw.New("0.0.0.0")
	if err != nil {
		panic(err)
	}

	p, err := New(r, "google.com")
	if err != nil {
		panic(err)
	}
	defer p.Close()

	for i := 0; i < 100; i++ {
		go func() {
			err := p.Send(1)
			if err != nil {
				panic(err)
			}

			_, value, err := p.Receive()
			if err != nil {
				panic(err)
			}

			spew.Dump(value.Bytes())
		}()
	}

	time.Sleep(time.Second * 10)
}
