package util

import (
	"log"

	scroll "github.com/meinside/scrollphat-go"
)

// Convert given 0/1 arrays to drawable bytes
func ArraysToBytes(arrays [][]uint8) []byte {
	if len(arrays) < scroll.NumRows || len(arrays[0]) < scroll.NumCols {
		log.Printf("wrong arrays length")
		return nil
	}

	bytes := []byte{}

	var b byte
	for col := 0; col < scroll.NumCols; col++ {
		b = 0
		for row := 0; row < scroll.NumRows; row++ {
			if arrays[row][col] > 0 {
				b |= (1 << uint(row))
			}
		}
		bytes = append(bytes, b)
	}

	return bytes
}
