package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var textureMap map[string]*sdl.Texture

// Init textureMap, called automatically
func init() {
	fmt.Println("INIT")
	textureMap = make(map[string]*sdl.Texture)
}

// LoadTexture loads texture to textureMap
func LoadTexture(file string, id string, r *sdl.Renderer) {
	image, err := img.Load(file)
	perror(err)
	defer image.Free()

	texture, err := r.CreateTextureFromSurface(image)
	perror(err)
	defer texture.Destroy()

	textureMap[id] = texture
}

// Draw draws game objects
func Draw(r *sdl.Renderer, gob *Gobject) {
	// var src, dest sdl.Rect
}

// DrawSprite draws sprite frame
func DrawSprite(r *sdl.Renderer, gob *Gobject) {
}

// TextureMapInfo prints textureMap
func TextureMapInfo() {
	fmt.Printf("%+v", textureMap)
}
