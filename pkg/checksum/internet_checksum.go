package checksum

import "bytes"

func CalculateInternetChecksum(payload *bytes.Buffer) uint16 {
	var checksum uint32
	var first, second uint16

	for {
		values := payload.Next(2)
		if len(values) == 0 {
			break
		} else if len(values) == 1 {
			first = uint16(values[0])
			second = 0
		} else {
			first = uint16(values[0])
			second = uint16(values[1])
		}

		currentWord := first | second<<8
		result := uint32(currentWord) + checksum
		checksum = result&0xFFFF + result>>16
	}

	return uint16(^checksum & 0xFFFF)
}
