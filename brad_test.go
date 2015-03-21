package brad

import (
	"math"
	"testing"
)

const tolerance = 0.00000000000001

func TestRadianConversion(t *testing.T) {
	var (
		b1 Brad    = 0
		r1 float64 = 0

		b2 Brad    = 1
		r2 float64 = math.Pi / 128

		b3 Brad    = 2
		r3 float64 = math.Pi / 64

		b4 Brad    = 4
		r4 float64 = math.Pi / 32

		b5 Brad    = 8
		r5 float64 = math.Pi / 16

		b6 Brad    = 16
		r6 float64 = math.Pi / 8

		b7 Brad    = 32
		r7 float64 = math.Pi / 4

		b8 Brad    = 192
		r8 float64 = math.Pi * 1.5

		b9 Brad    = 64
		r9 float64 = math.Pi / 2

		b10 Brad    = 128
		r10 float64 = math.Pi

		// 256 is all the way round, is the same as 0
		// and can't be represented in a uint8
		// 255  is one less than that
		b11 Brad    = 255
		r11 float64 = (2 * math.Pi) - (math.Pi / 128)
	)

	if b1.Radians() != r1 {
		t.Error(b1.Radians(), "!=", r1)
	}

	if b2.Radians() != r2 {
		t.Error(b2.Radians(), "!=", r2)
	}

	if b3.Radians() != r3 {
		t.Error(b3.Radians(), "!=", r3)
	}

	if b4.Radians() != r4 {
		t.Error(b4.Radians(), "!=", r4)
	}

	if b5.Radians() != r5 {
		t.Error(b5.Radians(), "!=", r5)
	}

	if b6.Radians() != r6 {
		t.Error(b6.Radians(), "!=", r6)
	}

	if b7.Radians() != r7 {
		t.Error(b7.Radians(), "!=", r7)
	}

	if b8.Radians() != r8 {
		t.Error(b8.Radians(), "!=", r8)
	}

	if b9.Radians() != r9 {
		t.Error(b9.Radians(), "!=", r9)
	}

	if b10.Radians() != r10 {
		t.Error(b10.Radians(), "!=", r10)
	}

	if b11.Radians() != r11 {
		t.Error(b11.Radians(), "!=", r11)
	}

}

// a very basic fuzzy equal, doesn't cover rounding
func equalish(a, b, tolerance float64) bool {
	if a == b {
		return true
	}

	diff := math.Abs(a - b)

	if diff <= tolerance {
		return true
	}

	return false
}
