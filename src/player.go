package main

import "github.com/veandco/go-sdl2/sdl"

// Player holds player data
type Player struct {
	Gobject
}

// Update updates player state
func (p *Gobject) Update() {
	keyStates := sdl.GetKeyboardState()
	p.speed = 4

	if keyStates[sdl.SCANCODE_DOWN] == 1 {
		p.fromX = 0
		p.currentFrame++
		if p.currentFrame >= p.framesY {
			p.currentFrame = 0
		}
		p.y += p.speed
		sdl.Delay(50)
	} else if keyStates[sdl.SCANCODE_UP] == 1 {
		p.fromX = 2 * p.width
		p.currentFrame++
		if p.currentFrame >= p.framesY {
			p.currentFrame = 0
		}
		p.y -= p.speed
		sdl.Delay(50)
	} else if keyStates[sdl.SCANCODE_LEFT] == 1 {
		p.fromX = p.width
		p.currentFrame++
		if p.currentFrame >= p.framesY {
			p.currentFrame = 0
		}
		p.x -= p.speed
		sdl.Delay(50)
	} else if keyStates[sdl.SCANCODE_RIGHT] == 1 {
		p.fromX = 3 * p.width
		p.currentFrame++
		if p.currentFrame >= p.framesY {
			p.currentFrame = 0
		}
		p.x += p.speed
		sdl.Delay(50)
	}

}

// Draw draws player
func (p *Gobject) Draw(r *sdl.Renderer) {
	p.fromY = p.currentFrame * p.width

	// Part of the spritesheet
	p.src = sdl.Rect{X: p.fromX, Y: p.fromY, W: p.width, H: p.height}
	// Part of the screen where to draw
	p.dest = sdl.Rect{X: p.x, Y: p.y, W: p.width, H: p.height}

	r.Copy(p.texture, &p.src, &p.dest)
}
