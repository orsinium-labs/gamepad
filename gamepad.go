package gamepad

import "github.com/0xcafed00d/joystick"

type GamePad struct {
	js joystick.Joystick
}

// Open the gamepad for reading, with the supplied id.
//
// Under linux the id is used to construct the joystick device name:
//   for example: id 0 will open device: "/dev/input/js0"
// Under Windows the id is the actual numeric id of the joystick
func NewGamepad(id int) (*GamePad, error) {
	js, err := joystick.Open(id)
	if err != nil {
		return nil, err
	}
	return &GamePad{js: js}, nil
}

// The current state of the gamepad
func (g *GamePad) State() (State, error) {
	jState, err := g.js.Read()
	state := State{
		axis:    append([]int(nil), jState.AxisData...),
		buttons: jState.Buttons,
	}
	return state, err
}
