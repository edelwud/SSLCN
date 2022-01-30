package main

import (
	"SSLCN/internal/ping"
	"SSLCN/internal/raw"
	"fmt"
	"sync"
)

func main() {
	r, err := raw.New("0.0.0.0")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	p, err := ping.New(r, "google.com")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func(i int) {
			dur, err := p.Run(uint16(i))
			if err != nil {
				panic(err)
			}

			fmt.Println(dur)
			wg.Done()
		}(i)
		wg.Add(1)
	}

	wg.Wait()
}
