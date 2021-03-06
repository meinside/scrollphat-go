package scroll

var TrimRedundantSpaces = true // trim redundant spaces in each character or not

// extracted from: https://github.com/pimoroni/scroll-phat/blob/master/library/font.png
var fontBytes [][]byte = [][]byte{
	{0x00, 0x00, 0x00, 0x00, 0x00}, // starts from 0x20
	{0x17, 0x00, 0x00, 0x00, 0x00},
	{0x03, 0x00, 0x03, 0x00, 0x00},
	{0x0A, 0x1F, 0x0A, 0x1F, 0x0A},
	{0x02, 0x15, 0x1F, 0x15, 0x08},
	{0x09, 0x04, 0x12, 0x00, 0x00},
	{0x0A, 0x15, 0x0A, 0x10, 0x00},
	{0x03, 0x00, 0x00, 0x00, 0x00},
	{0x0E, 0x11, 0x00, 0x00, 0x00},
	{0x11, 0x0E, 0x00, 0x00, 0x00},
	{0x14, 0x0E, 0x14, 0x00, 0x00},
	{0x04, 0x0E, 0x04, 0x00, 0x00},
	{0x10, 0x08, 0x00, 0x00, 0x00},
	{0x04, 0x04, 0x00, 0x00, 0x00},
	{0x10, 0x00, 0x00, 0x00, 0x00},
	{0x18, 0x06, 0x00, 0x00, 0x00},
	{0x0E, 0x15, 0x0E, 0x00, 0x00},
	{0x12, 0x1F, 0x10, 0x00, 0x00},
	{0x19, 0x15, 0x12, 0x00, 0x00},
	{0x11, 0x15, 0x0A, 0x00, 0x00},
	{0x0E, 0x09, 0x1C, 0x00, 0x00},
	{0x17, 0x15, 0x09, 0x00, 0x00},
	{0x0E, 0x15, 0x08, 0x00, 0x00},
	{0x19, 0x05, 0x03, 0x00, 0x00},
	{0x0A, 0x15, 0x0A, 0x00, 0x00},
	{0x02, 0x15, 0x0E, 0x00, 0x00},
	{0x0A, 0x00, 0x00, 0x00, 0x00},
	{0x10, 0x0A, 0x00, 0x00, 0x00},
	{0x04, 0x0A, 0x00, 0x00, 0x00},
	{0x0A, 0x0A, 0x0A, 0x00, 0x00},
	{0x0A, 0x04, 0x00, 0x00, 0x00},
	{0x01, 0x15, 0x02, 0x00, 0x00},
	{0x0E, 0x1D, 0x15, 0x0E, 0x00},
	{0x1E, 0x05, 0x1E, 0x00, 0x00},
	{0x1F, 0x15, 0x0A, 0x00, 0x00},
	{0x0E, 0x11, 0x11, 0x00, 0x00},
	{0x1F, 0x11, 0x0E, 0x00, 0x00},
	{0x1F, 0x15, 0x11, 0x00, 0x00},
	{0x1F, 0x05, 0x01, 0x00, 0x00},
	{0x0E, 0x11, 0x1D, 0x00, 0x00},
	{0x1F, 0x04, 0x1F, 0x00, 0x00},
	{0x11, 0x1F, 0x11, 0x00, 0x00},
	{0x09, 0x11, 0x0F, 0x00, 0x00},
	{0x1F, 0x04, 0x1B, 0x00, 0x00},
	{0x1F, 0x10, 0x10, 0x00, 0x00},
	{0x1F, 0x02, 0x04, 0x02, 0x1F},
	{0x1F, 0x02, 0x0C, 0x1F, 0x00},
	{0x0E, 0x11, 0x0E, 0x00, 0x00},
	{0x1F, 0x09, 0x06, 0x00, 0x00},
	{0x0E, 0x11, 0x09, 0x16, 0x00},
	{0x1F, 0x09, 0x16, 0x00, 0x00},
	{0x12, 0x15, 0x09, 0x00, 0x00},
	{0x01, 0x1F, 0x01, 0x00, 0x00},
	{0x0F, 0x10, 0x10, 0x0F, 0x00},
	{0x0F, 0x10, 0x0F, 0x00, 0x00},
	{0x0F, 0x10, 0x08, 0x10, 0x0F},
	{0x1B, 0x04, 0x1B, 0x00, 0x00},
	{0x03, 0x1C, 0x03, 0x00, 0x00},
	{0x19, 0x15, 0x13, 0x00, 0x00},
	{0x1F, 0x11, 0x00, 0x00, 0x00},
	{0x06, 0x18, 0x00, 0x00, 0x00},
	{0x11, 0x1F, 0x00, 0x00, 0x00},
	{0x02, 0x01, 0x02, 0x00, 0x00},
	{0x10, 0x10, 0x10, 0x00, 0x00},
	{0x01, 0x02, 0x00, 0x00, 0x00},
	{0x08, 0x14, 0x1C, 0x00, 0x00},
	{0x1F, 0x14, 0x08, 0x00, 0x00},
	{0x08, 0x14, 0x00, 0x00, 0x00},
	{0x08, 0x14, 0x1F, 0x00, 0x00},
	{0x0E, 0x15, 0x02, 0x00, 0x00},
	{0x1E, 0x05, 0x00, 0x00, 0x00},
	{0x02, 0x15, 0x0F, 0x00, 0x00},
	{0x1F, 0x04, 0x18, 0x00, 0x00},
	{0x1D, 0x00, 0x00, 0x00, 0x00},
	{0x10, 0x0D, 0x00, 0x00, 0x00},
	{0x1F, 0x04, 0x1A, 0x00, 0x00},
	{0x1F, 0x00, 0x00, 0x00, 0x00},
	{0x1C, 0x04, 0x18, 0x04, 0x18},
	{0x1C, 0x04, 0x18, 0x00, 0x00},
	{0x08, 0x14, 0x08, 0x00, 0x00},
	{0x1E, 0x0A, 0x04, 0x00, 0x00},
	{0x04, 0x0A, 0x1E, 0x00, 0x00},
	{0x1C, 0x02, 0x00, 0x00, 0x00},
	{0x14, 0x0A, 0x00, 0x00, 0x00},
	{0x0F, 0x12, 0x00, 0x00, 0x00},
	{0x0C, 0x10, 0x1C, 0x00, 0x00},
	{0x0C, 0x10, 0x0C, 0x00, 0x00},
	{0x0C, 0x10, 0x08, 0x10, 0x0C},
	{0x14, 0x08, 0x14, 0x00, 0x00},
	{0x02, 0x14, 0x0E, 0x00, 0x00},
	{0x12, 0x1A, 0x16, 0x00, 0x00},
	{0x04, 0x0E, 0x11, 0x00, 0x00},
	{0x1F, 0x00, 0x00, 0x00, 0x00},
	{0x11, 0x0E, 0x04, 0x00, 0x00},
	{0x08, 0x04, 0x08, 0x04, 0x00},
	{0x00, 0x00, 0x00, 0x00, 0x00},
}

// trim given bytes
func trim(bytes []byte) []byte {
	var checker byte = 0x00
	for _, b := range bytes {
		checker |= b
	}
	if checker == 0x00 { // all spaces
		return []byte{0x00, 0x00}
	}

	var left, right int
	// from left,
	for left = 0; left < len(bytes); left++ {
		if bytes[left] != 0x00 {
			break
		}
	}
	// from right
	for right = len(bytes) - 1; right >= 0; right-- {
		if bytes[right] != 0x00 {
			break
		}
	}

	return bytes[left : right+1]
}

// Get bytes array for given character
func BytesForCharacter(char rune) []byte {
	index := int(char) - 0x20 // font starts at 0x20

	// non-processable characters
	if index < 0 || index >= len(fontBytes) {
		index = 0 // default to '0x20'
	}

	bytes := fontBytes[index] // look up
	if TrimRedundantSpaces {
		bytes = trim(bytes)
	}

	return bytes
}

// Get bytes array for given string
func BytesForString(str string) []byte {
	bytes := []byte{}

	for _, char := range str {
		bytes = append(bytes, BytesForCharacter(char)...)
		bytes = append(bytes, 0x00) // 0x00 => 0x20(' ') = space between characters
	}

	return bytes
}
