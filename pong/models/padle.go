package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Padle struct {
	Position rl.Vector2
	Width    float32
	Height   float32
	Speed    float32
}

func NewPadle(Window rl.Rectangle, isPlayer bool) *Padle {
	padleWidth := float32(50)
	padleHeight := float32(300)
	initialPosition := rl.Vector2{}
	if isPlayer {
		initialPosition = rl.NewVector2(30, (Window.Height-padleHeight)/2)
	} else {
		initialPosition = rl.NewVector2(Window.Width-30-padleWidth, (Window.Height-padleHeight)/2)
	}
	return &Padle{
		Position: initialPosition,
		Width:    padleWidth,
		Height:   padleHeight,
		Speed:    600,
	}
}

func (p *Padle) Move(deltaTime float32, Window rl.Rectangle) {
	switch {
	case rl.IsKeyDown(rl.KeyW) && p.Position.Y > 0:
		p.Position.Y -= p.Speed * deltaTime
	case rl.IsKeyDown(rl.KeyS) && p.Position.Y+p.Height < Window.Height:
		p.Position.Y += p.Speed * deltaTime
	}
}

func (p *Padle) Bot(deltaTime float32, Window rl.Rectangle, bp rl.Vector2) {
	if bp.Y > p.Position.Y && bp.X > Window.Width/2 && p.Position.Y+p.Height < Window.Height {
		p.Position.Y += p.Speed * deltaTime
	} else if bp.Y < p.Position.Y && bp.X > Window.Width/2 && p.Position.Y > 0 {
		p.Position.Y -= p.Speed * deltaTime
	}
}

func (p *Padle) Draw() {
	rl.DrawRectangleV(p.Position, rl.NewVector2(p.Width, p.Height), rl.White)
}
