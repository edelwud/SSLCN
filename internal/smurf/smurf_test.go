package smurf

import "testing"

func TestSmurf_Run(t *testing.T) {
	smurf, err := New("127.0.0.1")
	if err != nil {
		panic(err)
	}

	err = smurf.Run()
	if err != nil {
		panic(err)
	}
}
