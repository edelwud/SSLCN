package checksum

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCalculateInternetChecksum(t *testing.T) {
	payload := []byte("Test payload\n")
	checksum := CalculateInternetChecksum(payload)
	fmt.Println(strconv.FormatInt(int64(checksum), 16))
}
