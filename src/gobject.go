package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

// GameObject interface
type GameObject interface {
	Draw(r *sdl.Renderer)
	Update()
}

// Gobject represents game object
type Gobject struct {
	// Image filename
	filename string
	// Key for mapping
	id string
	// Position
	x, y int32
	// Movement speed
	speed int32
	// Sprite size, not full image size
	width, height int32
	// For animation
	fromX, fromY, framesX, framesY, currentFrame, currentRow int32
	// Holds image
	texture *sdl.Texture
	// Part of the spritesheet
	src sdl.Rect
	// Part of the screen where to draw
	dest sdl.Rect
	// Is object moving
	isMoving bool
}

// NewGobject creates new game object
func NewGobject(r *sdl.Renderer, file, id string, x, y, framesX, framesY int32) *Gobject {
	gob := &Gobject{filename: file, x: x, y: y, framesX: framesX, framesY: framesY, speed: 1, currentFrame: 0}
	gob.Load(r)
	return gob
}

// Load texture
func (gob *Gobject) Load(r *sdl.Renderer) {
	image, err := img.Load(gob.filename)
	perror(err)
	defer image.Free()

	gob.texture, err = r.CreateTextureFromSurface(image)
	perror(err)

	// Query image size and calculate frame width and height
	_, _, imageWidth, imageHeight, _ := gob.texture.Query()
	gob.width = imageWidth / gob.framesX
	gob.height = imageHeight / gob.framesY
}

// Free resources
func (gob *Gobject) Free() {
	gob.texture.Destroy()
}
