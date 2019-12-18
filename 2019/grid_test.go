package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixedGrid(t *testing.T) {
	str := `#######
#.....#
#..^..#
#.<@>.#
#..V..#
#.....#
#######`
	g := NewFixedGrid(str)
	assert.Equal(t, string('#'), string(g.GetPoint(Point{0, 0})))
	assert.Equal(t, string('@'), string(g.GetPoint(Point{3, 3})))
	assert.Equal(t, str, g.String())
}
