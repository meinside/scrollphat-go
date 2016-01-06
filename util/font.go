package util

import (
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	fontFilename = "font.png" // 18 x 192 ([6 x 6] x [3 x 32])

	numCols = 3
	numRows = 32

	numFontPixels    = 5
	numFontGapPixels = 1
)

var fontPixels [numCols * numRows]FontPixels

type FontPixels struct {
	Pixels []byte
}

func checkPixel(color color.Color) byte {
	if r, g, b, _ := color.RGBA(); r == 0 && g == 0 && b == 0 {
		return 0x01
	}
	return 0x00
}

// Load fonts from font.png
func LoadFont() bool {
	if _, file, _, ok := runtime.Caller(0); ok {
		dir := filepath.Dir(file)
		fontFilepath := filepath.Join(dir, fontFilename)

		if fontFile, err := os.Open(fontFilepath); err == nil {
			defer fontFile.Close()

			if img, err := png.Decode(fontFile); err == nil {
				num := 0
				for cx := 0; cx < numCols; cx++ {
					for cy := 0; cy < numRows; cy++ {
						fontPixels[num].Pixels = make([]byte, numFontPixels)

						for y := 0; y < numFontPixels; y++ {
							for x := 0; x < numFontPixels; x++ {
								fontPixels[num].Pixels[x] |= (checkPixel(img.At(cx*(numFontPixels+numFontGapPixels)+x, cy*(numFontPixels+numFontGapPixels)+y)) << uint(y))
							}
						}

						num++
					}
				}

				return true
			}
		} else {
			log.Printf("font load failed: %s\n", err)
		}
	}

	return false
}

// Get bytes array for given character
func BytesForCharacter(char rune) []byte {
	index := int(char) - 0x20 // font starts at 0x20

	// non-processable characters
	if index < 0 || index >= len(fontPixels) {
		index = 0 // default to '0x20'
	}

	return fontPixels[index].Pixels
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
