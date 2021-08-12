package gamepad

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	is := require.New(t)
	g, err := NewGamepad(1)
	is.Nil(err)

	st, err := g.State()
	is.Nil(err)
	st.LS()

	// On linux, joystick lib reads the data asynchronously
	time.Sleep(100 * time.Millisecond)

	st, err = g.State()
	is.Nil(err)
	is.False(st.A())
	is.False(st.B())
	is.Equal(st.LT(), -100)
	is.Equal(st.LS().X, 0)
}

func Test_Equals(t *testing.T) {
	is := require.New(t)

	s := State{
		axis:    []int{0, 0},
		buttons: 0,
	}

	is.False(s.Equals(State{[]int{1, 0}, 0}))
	is.False(s.Equals(State{[]int{0, 1}, 0}))
	is.False(s.Equals(State{[]int{0, 0}, 1}))
	is.False(s.Equals(State{[]int{0, 0}, 1}))
	is.True(s.Equals(State{[]int{0, 0}, 0}))
}
