package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Window size
const (
	WindowWidth  = 640
	WindowHeight = 480
)

// Globals. TODO: Maybe Game struct
var (
	win       *sdl.Window
	rend      *sdl.Renderer
	keyStates []uint8
	err       error
	isRunning bool
)

// Error checker
func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err = sdl.Init(sdl.INIT_VIDEO)
	perror(err)

	win, err = sdl.CreateWindow(
		"Game",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		WindowWidth,
		WindowHeight,
		sdl.WINDOW_SHOWN,
	)
	perror(err)
	defer win.Destroy()

	rend, err = sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	perror(err)
	defer rend.Destroy()

	// Create player
	p1 := NewGobject(rend, "assets/george.png", 10, 10, 4, 4)
	fmt.Printf("%+v", p1)

	var event sdl.Event
	isRunning = true

	// Game loop
	for isRunning {

		// ORDER IS: events, update, draw

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				fmt.Println(t)
				isRunning = false
			}
		}

		p1.Update()
		p1.Draw(rend)

		sdl.Delay(16)

	} // End of isRunning

	p1.Free()
	sdl.Quit()
}
