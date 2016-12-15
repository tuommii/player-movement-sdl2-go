package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Window size
const (
	FPS          uint32 = 60
	DelayTime    uint32 = uint32(1000.0 / FPS)
	WindowWidth         = 640
	WindowHeight        = 480
	WindowTitle         = "Game"
)

// Globals
var (
	win       *sdl.Window
	rend      *sdl.Renderer
	event     sdl.Event
	err       error
	isRunning = true
)

// Error checker
func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Init SDL and create window
	err = sdl.Init(sdl.INIT_VIDEO)
	perror(err)

	win, err = sdl.CreateWindow(
		WindowTitle,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		WindowWidth,
		WindowHeight,
		sdl.WINDOW_SHOWN,
	)
	perror(err)
	defer win.Destroy()

	// Create renderer
	rend, err = sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	perror(err)
	defer rend.Destroy()

	// Load resources
	LoadTexture("assets/george.png", "player", rend)
	TextureMapInfo()

	// Game loop
	for isRunning {

		frameStartTime := sdl.GetTicks()

		// Handle event queue
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				fmt.Println(t)
				isRunning = false
			case *sdl.KeyDownEvent:
				OnKeyDown(event)
			case *sdl.KeyUpEvent:
				OnKeyUp(event)
			}
		}

		// Update
		// Render

		// If too fast add delay
		frameTime := sdl.GetTicks() - frameStartTime
		if frameTime < DelayTime {
			sdl.Delay(uint32(DelayTime - frameTime))
		}

	} // End of isRunning

	sdl.Quit()
}
