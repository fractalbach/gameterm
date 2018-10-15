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
)

const (
	bar = `=========================[ ~ ! gameterm ! ~ ]=========================`
)

var (
	player0    = '@'
	xmax, ymax = 30, 30
	xmin, ymin = 2, 2
	myX, myY   = 10, 10
	c1         = termbox.Attribute(47)
	bg         = termbox.Attribute(37)
)

func set(x, y int, r rune) {
	termbox.SetCell(x, y, r,
		termbox.ColorDefault, termbox.ColorDefault)
}

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

func box() {
	set(xmin-1, ymin-1, '╔')
	set(xmin-1, ymax+1, '╚')
	set(xmax+1, ymin-1, '╗')
	set(xmax+1, ymax+1, '╝')
	for i := xmin; i <= xmax; i++ {
		set(i, ymin-1, '═')
		set(i, ymax+1, '═')
	}
	for i := ymin; i <= ymax; i++ {
		set(xmin-1, i, '║')
		set(xmax+1, i, '║')
	}
}

func clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func draw() {
	clear()
	box()
	termbox.SetCell(myX, myY, player0, c1, bg)
	termbox.Flush()
}

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
		ev := termbox.PollEvent();
		
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
