package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Global consts
const (
	FPS          uint32 = 60
	DelayTime    uint32 = uint32(1000.0 / FPS)
	WindowWidth         = 640
	WindowHeight        = 480
	WindowTitle         = "Game"
)

// Globals, maybe someday wrapped to struct but now less to type
var (
	win       *sdl.Window
	rend      *sdl.Renderer
	event     sdl.Event
	err       error
	isRunning = true
	// Maybe later
	gameObjects map[string]*Gobject
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

	// Create player
	player := NewGobject(rend, "assets/george.png", "player", 10, 10, 4, 4)
	// Init gameObjects map
	gameObjects = make(map[string]*Gobject)
	gameObjects[player.id] = player

	// Game loop
	for isRunning {

		frameStartTime := sdl.GetTicks()

		// Handle event queue
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				fmt.Println(t)
				isRunning = false
			}
		}

		// Clear screen
		rend.SetDrawColor(0, 255, 0, 255)
		rend.Clear()

		// Update
		player.Update()

		// Render
		player.Draw(rend)

		rend.Present()

		// If too fast add delay
		frameTime := sdl.GetTicks() - frameStartTime
		if frameTime < DelayTime {
			sdl.Delay(uint32(DelayTime - frameTime))
		}

	} // End of isRunning

	player.Free()
	sdl.Quit()
}
