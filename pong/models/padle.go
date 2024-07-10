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
	padleWidth := float32(20)
	padleHeight := float32(100)
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
		Speed:    500,
	}
}

func (p *Padle) Draw() {
	rl.DrawRectangleV(p.Position, rl.NewVector2(p.Width, p.Height), rl.White)
}
