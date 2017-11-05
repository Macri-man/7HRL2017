package main

import (
	"fmt"
	"math/rand"
	"time"

	termbox "github.com/nsf/termbox-go"
)

type cell int

const (
	empty cell = iota
	wall
	ground
	doorh
	doorw
	potion
	creature
)

func (char cell) DrawCell() rune {
	switch char {
	case empty:
		return '.'
	case wall:
		return '#'
	case doorh:
		return '-'
	case doorw:
		return '|'
	case ground:
		return '*'
	case potion:
		return 'P'
	case creature:
		return 'C'
	}
	return '!'
}

func (char cell) Walkable() bool {
	walkable := true
	if char == wall {
		walkable = false
	}
	return walkable
}

type Room struct {
	id string
	x  int
	y  int
	w  int
	h  int
}

type Item struct {
	id     string
	x      int
	y      int
	effect func() (string, int)
}

type Creature struct {
	id     string
	x      int
	y      int
	attack func() (string, int)
	effect func(num int) (string, int)
}

type Display struct {
	width     int
	height    int
	contents  []cell
	rooms     []Room
	creatures []Creature
	items     []Item
}

func (d Display) AT(x, y int) cell {
	return d.contents[x+d.width*y]
}

func (d Display) ATPlayer(x, y int) cell {
	return d.contents[x+d.width*y]
}

func (d *Display) SET(x, y int, char cell) {
	d.contents[x+d.width*y] = char
	//fmt.Println(fmt.Sprintf("w %d h %d w*h %d", d.width, d.height, d.width*d.height))
	//fmt.Println(fmt.Sprintf("x %d d.height %d y %d", x, d.width, y))
	//fmt.Println(x + d.width*y)
}

func makeDisplay(w, h, numOfRooms int) Display {

	contents := make([]cell, w*h)

	items := make([]Item, 0)
	creatures := make([]Creature, 0)

	//startx := 0 // w/2
	//starty := 0 // h/2
	rangeL := 2
	rangeH := 6

	rooms := make([]Room, 0)

	//rooms = append(rooms, makeRoom(i, startx, starty, rangeR, rangeR))

	for i := 0; i < numOfRooms; i++ {
		rooms = append(rooms, makeRoom(i, getRange(3, w-2), getRange(3, h-2), rangeL, rangeH))
	}

	return Display{w, h, contents, rooms, creatures, items}
}

func (d *Display) setcontents() {
	for _, r := range d.rooms {
		for index := r.y; index < r.y+r.w; index++ {
			d.SET(index, r.y, wall)
			d.SET(index, r.h-1, wall)
		}

		for index := r.x; index < r.x+r.h; index++ {
			d.SET(r.y, index, wall)
			d.SET(r.w-1, index, wall)
		}
	}
}

func getRange(low, high int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(high-low) + low
}

func makeRoom(id, x, y, rangeL, rangeH int) Room {
	return Room{fmt.Sprintf("Room %d", id), x, y, getRange(rangeL, rangeH), getRange(rangeL, rangeH)}
}

func makeHallway(id, room1, room2, x, y, rangeL, rangeH int) Room {
	return Room{fmt.Sprintf("Hallway %d connecting Room %d Room %d", id, room1, room2), x, y, getRange(rangeL, rangeH), getRange(rangeL, rangeH)}
}

func (d Display) DisplayScreen(xPlayer, yPlayer int) {

	w, h := termbox.Size()
	//tempx := xPlayer + (w / 2)
	//tempy := yPlayer + (h / 2)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			termbox.SetCell(x, y, d.AT(x, y).DrawCell(), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.Flush()
}
