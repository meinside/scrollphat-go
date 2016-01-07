// Go library for Pimoroni's Scroll pHat (Rasbperry Pi)

package scroll

import (
	"log"
	"sync"

	i2c "github.com/meinside/scrollphat-go/i2c"
)

const (
	i2cAddr  = 0x60
	mode5x11 = 0x03

	// commands
	CmdSetMode       = 0x00
	CmdSetBrightness = 0x19
	CmdPaintPixel    = 0x01

	// 11 x 5 LEDs
	NumRows = 5
	NumCols = 11
)

type ScrollPHat struct {
	sync.Mutex
	i2c *i2c.I2C

	// internal pixels array
	pixels []byte
}

// Get a new ScrollPHat
func New() *ScrollPHat {
	if i, err := i2c.New(i2cAddr); err == nil {
		nu := ScrollPHat{
			i2c:    i,
			pixels: make([]byte, NumCols),
		}

		if _, err := nu.write(CmdSetMode, []byte{mode5x11}); err != nil {
			log.Printf("*** failed to set mode: %s\n", err)
		}

		return &nu
	} else {
		log.Printf("*** failed to open i2c: %s\n", err)
	}

	return nil
}

// Write given command and data to i2c
func (s ScrollPHat) write(cmd byte, data []byte) (n int, err error) {
	var bytes []byte = []byte{cmd}
	bytes = append(bytes, data...)

	return s.i2c.Write(bytes)
}

// Set LED brightness
func (s ScrollPHat) SetBrightness(brightness byte) {
	s.Lock()
	defer s.Unlock()

	if _, err := s.write(CmdSetBrightness, []byte{brightness}); err != nil {
		log.Printf("[SetBrightness] error writing to i2c: %s\n", err)
	}
}

// Draw given pixel bytes
func (s ScrollPHat) draw(bytes []byte) {
	if bytes == nil {
		return
	}

	s.Lock()
	defer s.Unlock()

	pixelBytes := append(bytes, 0xFF) // need trailing 0xFF for drawing

	if _, err := s.write(CmdPaintPixel, pixelBytes); err != nil {
		log.Printf("[draw] error writing to i2c: %s\n", err)
	}
}

// Alter pixels in the memory and draw them
func (s ScrollPHat) DrawBytes(bytes []byte) {
	copy(s.pixels, bytes)

	s.draw(s.pixels)
}

// Alter given pixel in the memory and draw them
func (s ScrollPHat) DrawPixel(col, row uint, isOn bool) {
	if isOn {
		s.pixels[col] |= (1 << row)
	} else {
		s.pixels[col] ^= (1 << row)
	}

	s.draw(s.pixels)
}

// Toggle given pixel in the memory and draw them
func (s ScrollPHat) TogglePixel(col, row uint) {
	if s.pixels[col]&(1<<row) > 0 {
		s.pixels[col] ^= (1 << row)
	} else {
		s.pixels[col] |= (1 << row)
	}

	s.draw(s.pixels)
}

// Clear pixels for both memory and LED
func (s ScrollPHat) Clear() {
	for i := 0; i < NumCols; i++ {
		s.pixels[i] = 0x00
	}

	s.draw(s.pixels)
}

// Fill pixels for both memory and LED
func (s ScrollPHat) Fill() {
	var fill byte
	for i := 0; i < NumRows; i++ {
		fill |= (1 << uint(i))
	}

	for i := 0; i < NumCols; i++ {
		s.pixels[i] = fill
	}

	s.draw(s.pixels)
}

// Invert pixels for both memory and LED
func (s ScrollPHat) Invert() {
	for i := 0; i < NumCols; i++ {
		for j := 0; j < NumRows; j++ {
			if s.pixels[i]&(1<<uint(j)) > 0 {
				s.pixels[i] ^= (1 << uint(j))
			} else {
				s.pixels[i] |= (1 << uint(j))
			}
		}
	}

	s.draw(s.pixels)
}

// Close
func (s ScrollPHat) Close() error {
	return s.i2c.Close()
}
