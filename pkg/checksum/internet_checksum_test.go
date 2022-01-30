package checksum

import (
	"bytes"
	"strconv"
	"testing"
)

func TestCalculateInternetChecksum(t *testing.T) {
	payload := new(bytes.Buffer)
	payload.Write([]byte("Test payload\n"))

	checksum := CalculateInternetChecksum(payload)

	if strconv.FormatInt(int64(checksum), 16) != "68de" {
		t.Fatal("incorrect internet checksum result")
	}
}
