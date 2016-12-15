package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var keyboardState []uint8

// OnKeyDown is called when key is pressed
func OnKeyDown(ev sdl.Event) {
	keyboardState = sdl.GetKeyboardState()
	fmt.Println("DOWN BUTTON PRESSED:", IsKeyPressed(sdl.SCANCODE_DOWN))
	fmt.Printf("%+v\n", ev)
}

// OnKeyUp is called when key is released
func OnKeyUp(ev sdl.Event) {
	keyboardState = nil
}

// IsKeyPressed return true if key is pressed
func IsKeyPressed(key sdl.Scancode) bool {
	if len(keyboardState) != 0 {
		if keyboardState[key] == 1 {
			return true
		}
	}
	return false
}
