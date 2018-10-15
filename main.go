/*
======================================================================
			 Messing with Termbox
======================================================================

Experiments with Termbox for grids and stuff like that.  Might Turns
this into a game interface later on.


Some symbols
--------------------
웃
ጰ
♘
♀
ϕ
'♀' | '̪'

More symbols
--------------------
https://en.wikibooks.org/wiki/Unicode/List_of_useful_symbols


*/

package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
	"golang.org/x/text/width"
)

const (
	bar = `=========================[ ~ ! gameterm ! ~ ]=========================`
)

var (
	xmax, ymax = 30, 30
	xmin, ymin = 2, 2
	myX, myY   = 10, 10
	c1         = termbox.Attribute(47)
	bg         = termbox.Attribute(37)
)

var (
	player0   = '@'
	wallNW    = widen('╔')
	wallNE    = widen('╗')
	wallSW    = widen('╚')
	wallSE    = widen('╝')
	wallHoriz = widen('═')
	wallVert  = widen('║')
)

// Controls
// ----------------------------------------

func up() {
	if myY > ymin {
		myY--
	}
}
func down() {
	if myY < ymax {
		myY++
	}
}
func left() {
	if myX > xmin {
		myX--
	}
}
func right() {
	if myX < xmax {
		myX++
	}
}

// Display Backgrounds
// ----------------------------------------

func box() {
	set(xmin-1, ymin-1, wallNW)
	set(xmin-1, ymax+1, wallSW)
	set(xmax+1, ymin-1, wallNE)
	set(xmax+1, ymax+1, wallSE)
	for i := xmin; i <= xmax; i++ {
		set(i, ymin-1, wallHoriz)
		set(i, ymax+1, wallHoriz)
	}
	for i := ymin; i <= ymax; i++ {
		set(xmin-1, i, wallVert)
		set(xmax+1, i, wallVert)
	}
}

func draw() {
	clear()
	box()
	termbox.SetCell(myX, myY, player0, c1, bg)
	termbox.Flush()
}

// Helper functions
// ----------------------------------------

func widen(r rune) rune {
	w := width.Widen.String(string(r))
	var wr rune
	for _, v := range w {
		wr = v
		break
	}
	return wr
}

func set(x, y int, r rune) {
	termbox.SetCell(x, y, r,
		termbox.ColorDefault, termbox.ColorDefault)
}

func clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

// Main Loop
// ----------------------------------------

func enterTerm() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.OutputNormal)

loop:
	for {
		draw()
		ev := termbox.PollEvent()

		switch ev.Ch {
		case 'q':
			break loop
		}

		switch ev.Key {
		case termbox.KeyEsc:
			break loop
		case termbox.KeySpace:
			break loop
		case termbox.KeyArrowDown:
			down()
		case termbox.KeyArrowLeft:
			left()
		case termbox.KeyArrowRight:
			right()
		case termbox.KeyArrowUp:
			up()
		}
	}
}

func main() {
	startTime := time.Now()
	enterTerm()
	fmt.Println("Started gameterm at: ", startTime)
	fmt.Println("Stoped gameterm at:", time.Now().Sub(startTime))
}
