// Package oshift implements the origin shift algorithm laid out in
// https://youtube.com/???? and
// https://github.com/CaptainLuma/New-Maze-Generating-Algorithm
package oshift

import (
	"math/rand"
)

// New returns a new Maze (TODO: initialize). Expect undefined behavior for
// widths or heights less than 2.
func New(width, height int) Maze {
	grid := make([][]*node, width)
	for i := range grid {
		grid[i] = make([]*node, height)
		for j := range grid[i] {
			grid[i][j] = &node{i, j, nil}
			grid[i][j].edge = grid[i][j]
		}
	}
	return Maze{
		grid:   grid,
		origin: nil,
		width:  width,
		height: height,
	}
}

type Maze struct {
	grid          [][]*node
	origin        *node
	width, height int
}

func (m Maze) Width() int  { return m.Width() }
func (m Maze) Height() int { return m.Height() }

func (m Maze) Shuffle(iter int) {
	for i := 0; i < iter; i++ {
		var dx, dy int
		switch rand.Int() % 4 {
		case 0:
			if m.origin.x == 0 {
				dx, dy = 1, 0
			}
			dx, dy = -1, 0
		case 1:
			if m.origin.x == m.Width()-1 {
				dx, dy = -1, 0
			}
			dx, dy = 1, 0
		case 2:
			if m.origin.y == 0 {
				dx, dy = 0, 1
			}
			dx, dy = 0, -1
		case 3:
			if m.origin.y == m.Height()-1 {
				dx, dy = 0, -1
			}
			dx, dy = 0, 1
		default:
			panic("THE END IS NIGH! REALITY HAS BROKEN!")
		}
		m.origin.edge = m.grid[m.origin.x+dx][m.origin.y+dy]
		m.origin = m.origin.edge
		m.origin.edge = m.origin.edge
	}
}

// Path returns (-1, -1) if the query is out of bounds, and (x, y) if it's the
// origin
func (m Maze) Follow(x, y int) (int, int) {
	if x < 0 || x >= m.Width() || y < 0 || y >= m.Height() {
		return -1, -1
	}
	return m.grid[x][y].edge.x, m.grid[x][y].edge.y
}

func (m Maze) Verify() bool {
	checked := make([][]bool, m.Width())
	for i := range checked {
		checked[i] = make([]bool, m.Height())
	}
	for i := range checked {
	EACH:
		for j := range checked {
			if checked[i][j] {
				continue
			}
			node := m.grid[i][j]
			for path := 0; path < m.width*m.height; path++ {
				checked[node.x][node.y] = true
				if node == m.origin {
					continue EACH
				}
				if node.edge == node {
					return false
				}
				node = node.edge
			}
			return false
		}
	}
	return false
}
