package traceroute

import (
	"SSLCN/internal/raw"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestTraceroute_Run(t *testing.T) {
	r, err := raw.New("0.0.0.0")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	traceroute, err := New(r, "google.com")
	if err != nil {
		panic(err)
	}

	routes, err := traceroute.Run(64)
	if err != nil {
		panic(err)
	}

	spew.Dump(routes)
}
