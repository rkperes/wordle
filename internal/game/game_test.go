package game

import (
	"testing"

	"github.com/rafaelkperes/wordle/internal/dict"
	"github.com/stretchr/testify/require"
)

func TestGameTyping(t *testing.T) {
	rq := require.New(t)

	d, err := dict.NewLocalDict()
	rq.NoError(err)
	g := NewGame(d)

	// number of tries not relevant for this test
	g.Start(1)
	rq.Equal([]string{""}, g.Board())

	g.Type("a")
	rq.Equal([]string{"a"}, g.Board())
	g.Type("b")
	rq.Equal([]string{"ab"}, g.Board())
	g.Backspace()
	rq.Equal([]string{"a"}, g.Board())

	g.Type("b")
	g.Type("b")
	g.Type("b")
	g.Type("b")
	rq.Equal([]string{"abbbb"}, g.Board())
	g.Type("c")
	rq.Equal([]string{"abbbb"}, g.Board())

	g.Backspace()
	g.Backspace()
	g.Backspace()
	g.Backspace()
	rq.Equal([]string{"a"}, g.Board())
	g.Backspace()
	g.Backspace()
	rq.Equal([]string{""}, g.Board())
}
