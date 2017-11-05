package main

import termbox "github.com/nsf/termbox-go"

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	d := makeDisplay(100, 100, 10)
	d.setcontents()

mainloop:
	for {
		d.DisplayScreen(10, 10)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			case termbox.KeyArrowRight:
				//move player right
			case termbox.KeyArrowLeft:
				//move player left
			case termbox.KeyArrowUp:
				//move player up
			case termbox.KeyArrowDown:
				//move player down
			default:

			}
		case termbox.EventResize:
			d.width = ev.Width
			d.height = ev.Height
		case termbox.EventError:
			panic(ev.Err)
		}
		d.DisplayScreen(10, 10)
	}

}
