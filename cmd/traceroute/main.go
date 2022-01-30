package main

import (
	"SSLCN/internal/raw"
	"SSLCN/internal/traceroute"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	r, err := raw.New("0.0.0.0")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	tr, err := traceroute.New(r, "google.com")
	if err != nil {
		panic(err)
	}

	routes, err := tr.Run(64)
	if err != nil {
		panic(err)
	}

	spew.Dump(routes)
}
