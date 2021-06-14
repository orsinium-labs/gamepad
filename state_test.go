package gamepad_test

import (
	"testing"
	"time"

	"github.com/orsinium-labs/gamepad"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	is := require.New(t)
	g, err := gamepad.NewGamepad(1)
	is.Nil(err)

	// On linux, joystick lib reads the data asynchronously
	time.Sleep(100 * time.Millisecond)

	st, err := g.State()
	is.Nil(err)
	is.False(st.A())
	is.Equal(st.LT(), -100)
}