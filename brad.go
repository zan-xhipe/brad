package brad

import "math"

type Brad uint8

const (
	radianScale = math.Pi / 128
)

// Radians returns the angle of the Brad in radians
func (b Brad) Radians() float64 {
	return radianScale * float64(b)
}
