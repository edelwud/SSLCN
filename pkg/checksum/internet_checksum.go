package checksum

func CalculateInternetChecksum(payload []byte) uint32 {
	var checksum uint32
	for i := 0; i < len(payload); i += 2 {
		var first, second uint16
		if i+1 == len(payload) {
			second = 0
		} else {
			second = uint16(payload[i+1])
		}

		first = uint16(payload[i])

		currentWord := first<<8 | second
		result := uint32(currentWord) + checksum
		checksum = result&0xFFFF + result>>16
	}
	return ^checksum & 0xFFFF
}
