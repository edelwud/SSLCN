package ping

import (
	"SSLCN/internal/raw"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	r, err := raw.New("0.0.0.0")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	_, err = New(r, "google.com")
	if err != nil {
		panic(err)
	}
}

func TestPing_Send(t *testing.T) {
	r, err := raw.New("0.0.0.0")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	p, err := New(r, "1.1.1.1")
	if err != nil {
		panic(err)
	}

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
	defer r.Close()

	p, err := New(r, "1.1.1.1")
	if err != nil {
		panic(err)
	}

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
	defer r.Close()

	p, err := New(r, "google.com")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func(i int) {
			dur, err := p.Run(uint16(i))
			if err != nil {
				return
			}
			spew.Dump(dur)
			wg.Done()
		}(i)
		wg.Add(1)
	}

	wg.Wait()
}
