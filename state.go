package gamepad

import (
	"image"
)

const (
	axisMod   = 327
	axisCount = 8
)

// State holds the state of the gamepad buttons.
//
// See the gist for buttons names and their placement on the gamepad:
//   https://gist.github.com/palmerj/586375bcc5bc83ccdaf00c6f5f863e86
type State struct {
	axis    []int
	buttons uint32
}

// Equals verify if actual state is equals to another
func (s State) Equals(state State) bool {
	// not verifying axis length (assuming it wont change)
	for i, ax := range s.axis {
		if ax != state.axis[i] {
			return false
		}
	}

	return s.buttons == state.buttons
}

// A button is pressed
func (s State) A() bool {
	return s.buttons&0b_0001 > 0
}

// B button is pressed
func (s State) B() bool {
	return s.buttons&0b_0010 > 0
}

// X button is pressed
func (s State) X() bool {
	return s.buttons&0b_0100 > 0
}

// Y button is pressed
func (s State) Y() bool {
	return s.buttons&0b_1000 > 0
}

// Left Bumper is pressed
func (s State) LB() bool {
	return s.buttons&0b_0001_0000 > 0
}

// Right Bumper is pressed
func (s State) RB() bool {
	return s.buttons&0b_0010_0000 > 0
}

// Back button is pressed
func (s State) Back() bool {
	return s.buttons&0b_0100_0000 > 0
}

// Start button is pressed
func (s State) Start() bool {
	return s.buttons&0b_1000_0000 > 0
}

// Guide button is pressed
func (s State) Guide() bool {
	return s.buttons&0b_0001_0000_0000 > 0
}

// Left Stick Button is pressed
func (s State) LSB() bool {
	return s.buttons&0b_0010_0000_0000 > 0
}

// Right Stick Button is pressed
func (s State) RSB() bool {
	return s.buttons&0b_0100_0000_0000 > 0
}

// Left Stick coordinates, from -100 to 100.
// Up position is negative.
func (s State) LS() image.Point {
	if len(s.axis) < axisCount {
		return image.Point{}
	}
	return image.Point{
		X: s.axis[0] / axisMod,
		Y: s.axis[1] / axisMod,
	}
}

// Right Stick coordinates, from -100 to 100.
// Up position is negative.
func (s State) RS() image.Point {
	if len(s.axis) < axisCount {
		return image.Point{}
	}
	return image.Point{
		X: s.axis[3] / axisMod,
		Y: s.axis[4] / axisMod,
	}
}

// Directional Pad left button is pressed
func (s State) DPadLeft() bool {
	if len(s.axis) < axisCount {
		return false
	}
	return s.axis[6] < 0
}

// Directional Pad right button is pressed
func (s State) DPadRight() bool {
	if len(s.axis) < axisCount {
		return false
	}
	return s.axis[6] > 0
}

// Directional Pad up button is pressed
func (s State) DPadUp() bool {
	if len(s.axis) < axisCount {
		return false
	}
	return s.axis[7] < 0
}

// Directional Pad down button is pressed
func (s State) DPadDown() bool {
	if len(s.axis) < axisCount {
		return false
	}
	return s.axis[7] > 0
}

// Left Trigger state from -100 to 100 where 100 is pressed.
func (s State) LT() int {
	if len(s.axis) < axisCount {
		return -100
	}
	return s.axis[2] / axisMod
}

// Right Trigger state from -100 to 100 where 100 is pressed.
func (s State) RT() int {
	if len(s.axis) < axisCount {
		return -100
	}
	return s.axis[5] / axisMod
}
