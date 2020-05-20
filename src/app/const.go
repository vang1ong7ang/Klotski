package main

var (
	// SHAPE ...
	SHAPE = [4]struct {
		W int
		H int
	}{
		{2, 2},
		{2, 1},
		{1, 2},
		{1, 1},
	}
	// DIRECTS ...
	DIRECTS = []struct {
		X int
		Y int
	}{
		{-1, 0},
		{+1, 0},
		{0, -1},
		{0, +1},
	}
)

const (
	// SIZEX ...
	SIZEX = 4
	// SIZEY ...
	SIZEY = 5
	// MAXSTEP ...
	MAXSTEP = 0x1000
)
