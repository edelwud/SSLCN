package checksum

import (
	"strconv"
	"testing"
)

func TestCalculateInternetChecksum(t *testing.T) {
	payload := []byte("Test payload\n")
	checksum := CalculateInternetChecksum(payload)
	if strconv.FormatInt(int64(checksum), 16) != "de68" {
		t.Fatal("incorrect internet checksum result")
	}
}
