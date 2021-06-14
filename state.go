package gamepad

import (
	"image"

	"github.com/0xcafed00d/joystick"
)

const axisMod = 327

// State holds the state of the gamepad buttons.
//
// See the gist for buttons names and their placement on the gamepad:
//   https://gist.github.com/palmerj/586375bcc5bc83ccdaf00c6f5f863e86
type State struct {
	state joystick.State
}

// A button is pressed
func (s *State) A() bool {
	return s.state.Buttons&0b_0001 == 1
}

// B button is pressed
func (s *State) B() bool {
	return s.state.Buttons&0b_0010 == 1
}

// X button is pressed
func (s *State) X() bool {
	return s.state.Buttons&0b_0100 == 1
}

// Y button is pressed
func (s *State) Y() bool {
	return s.state.Buttons&0b_1000 == 1
}

// Left Bumper is pressed
func (s *State) LB() bool {
	return s.state.Buttons&0b_0001_0000 == 1
}

// Right Bumper is pressed
func (s *State) RB() bool {
	return s.state.Buttons&0b_0010_0000 == 1
}

// Back button is pressed
func (s *State) Back() bool {
	return s.state.Buttons&0b_0100_0000 == 1
}

// Start button is pressed
func (s *State) Start() bool {
	return s.state.Buttons&0b_1000_0000 == 1
}

// Guide button is pressed
func (s *State) Guide() bool {
	return s.state.Buttons&0b_0001_0000_0000 == 1
}

// Left Stick Button is pressed
func (s *State) LSB() bool {
	return s.state.Buttons&0b_0010_0000_0000 == 1
}

// Right Stick Button is pressed
func (s *State) RSB() bool {
	return s.state.Buttons&0b_0100_0000_0000 == 1
}

// Left Stick coordinates, from -100 to 100.
// Up position is negative.
func (s *State) LS() image.Point {
	return image.Point{
		X: s.state.AxisData[0] / axisMod,
		Y: s.state.AxisData[1] / axisMod,
	}
}

// Right Stick coordinates, from -100 to 100.
// Up position is negative.
func (s *State) RS() image.Point {
	return image.Point{
		X: s.state.AxisData[3] / axisMod,
		Y: s.state.AxisData[4] / axisMod,
	}
}

// Directional Pad left button is pressed
func (s *State) DPadLeft() bool {
	return s.state.AxisData[6] < 0
}

// Directional Pad right button is pressed
func (s *State) DPadRight() bool {
	return s.state.AxisData[6] > 0
}

// Directional Pad up button is pressed
func (s *State) DPadUp() bool {
	return s.state.AxisData[7] < 0
}

// Directional Pad down button is pressed
func (s *State) DPadDown() bool {
	return s.state.AxisData[7] > 0
}

// Left Trigger state from -100 to 100 where 100 is pressed.
func (s *State) LT() int {
	return s.state.AxisData[2] / axisMod
}

// Right Trigger state from -100 to 100 where 100 is pressed.
func (s *State) RT() int {
	return s.state.AxisData[5] / axisMod
}
