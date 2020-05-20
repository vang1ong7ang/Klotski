package main

import (
	"fmt"
	"sort"
)

// State ...
type State [10][3]int

// step ...
func (state State) step(n int, x int, y int) (ret State, ok bool) {
	S, X, Y := state[n][0], state[n][1], state[n][2]
	W, H := SHAPE[S].W, SHAPE[S].H
	XL, YL := X+x, Y+y
	XH, YH := XL+W, YL+H
	if XL < 0 {
		return
	}
	if YL < 0 {
		return
	}
	if XH > SIZEX {
		return
	}
	if YH > SIZEY {
		return
	}

	for i := range state {
		if i == n {
			continue
		}
		SI, XIL, YIL := state[i][0], state[i][1], state[i][2]
		WI, HI := SHAPE[SI].W, SHAPE[SI].H
		XIH, YIH := XIL+WI, YIL+HI
		if XL >= XIH {
			continue
		}
		if XIL >= XH {
			continue
		}
		if YL >= YIH {
			continue
		}
		if YIL >= YH {
			continue
		}
		return
	}

	ok = true
	for i := range state {
		if i == n {
			ret[i][0] = S
			ret[i][1] = XL
			ret[i][2] = YL
			continue
		}
		ret[i] = state[i]
	}
	return
}

// Move ...
func (state State) Move(n int, x int, y int) (ret State, ok bool) {
	ret, ok = state.step(n, x, y)
	if ok == false {
		return
	}
	sort.Slice(ret[:], func(i int, j int) bool {
		for k := 0; k < 3; k++ {
			if ret[i][k] < ret[j][k] {
				return true
			}
			if ret[i][k] > ret[j][k] {
				return false
			}
		}
		return false
	})
	return
}

// Move2 ...
func (state State) Move2(n int, x1 int, y1 int, x2 int, y2 int) (ret State, ok bool) {
	ret, ok = state.step(n, x1, y1)
	if ok == false {
		return
	}
	ret, ok = ret.step(n, x2, y2)
	if ok == false {
		return
	}
	sort.Slice(ret[:], func(i int, j int) bool {
		for k := 0; k < 3; k++ {
			if ret[i][k] < ret[j][k] {
				return true
			}
			if ret[i][k] > ret[j][k] {
				return false
			}
		}
		return false
	})
	return
}

// Moves ...
func (state State) Moves() (ret []State) {
	for i := range state {
		for _, d := range DIRECTS {
			t, ok := state.Move(i, d.X, d.Y)
			if ok == false {
				continue
			}
			ret = append(ret, t)
		}
	}

	for i := range state {
		for _, d := range DIRECTS {
			for _, dd := range DIRECTS {
				if d.X+dd.X == 0 && d.Y+dd.Y == 0 {
					continue
				}
				t, ok := state.Move2(i, d.X, d.Y, dd.X, dd.Y)
				if ok == false {
					continue
				}
				ret = append(ret, t)
			}
		}
	}
	return
}

// Final ...
func (state State) Final() bool {
	if state[0][0] != 0 {
		return false
	}
	if state[0][1] != 1 {
		return false
	}
	if state[0][2] != 3 {
		return false
	}
	return true
}

// Print ...
func (state State) Print() {
	board := [4][5]interface{}{}
	for i := range state {
		S, X, Y := state[i][0], state[i][1], state[i][2]
		W, H := SHAPE[S].W, SHAPE[S].H
		for x := X; x < X+W; x++ {
			for y := Y; y < Y+H; y++ {
				board[x][y] = i
			}
		}
	}
	fmt.Println(state)
	for y := 0; y < SIZEY; y++ {
		for x := 0; x < SIZEX; x++ {
			if board[x][y] != nil {
				fmt.Print(board[x][y])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
