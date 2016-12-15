package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var keyBoardState []uint8

// OnKeyDown is called when key is pressed
func OnKeyDown(ev sdl.Event) {
	keyBoardState = sdl.GetKeyboardState()
	fmt.Println("DOWN BUTTON PRESSED:", IsKeyPressed(sdl.SCANCODE_DOWN))
	fmt.Printf("%+v\n", ev)
}

// OnKeyUp is called when key is released
func OnKeyUp(ev sdl.Event) {
	keyBoardState = nil
}

// IsKeyPressed return true if key is pressed
func IsKeyPressed(key sdl.Scancode) bool {
	if len(keyBoardState) != 0 {
		if keyBoardState[key] == 1 {
			return true
		}
	}
	return false
}
