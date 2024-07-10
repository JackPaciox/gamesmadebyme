package models

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Position     rl.Vector2
	Radius       float32
	InitialSpeed float32
	acceletation float32
	Direction    rl.Vector2
}

func NewBall(Window rl.Rectangle) *Ball {

	directionX := rand.Float32() * Window.Width
	directionY := rand.Float32() * Window.Height

	InitialDirection := rl.NewVector2(directionX, directionY)
	InitialPosition := rl.NewVector2(Window.Width/2, Window.Height/2)

	return &Ball{
		Position:     InitialPosition,
		Radius:       20,
		InitialSpeed: 500,
		acceletation: 1.15,
		Direction:    InitialDirection,
	}
}

func (b *Ball) Reset(Window rl.Rectangle) {
	b.Position = rl.NewVector2(Window.Width/2, Window.Height/2)
	b.Direction = rl.NewVector2(rand.Float32()*Window.Width, rand.Float32()*Window.Height)
}

func (b *Ball) Draw() {
	rl.DrawCircleV(b.Position, b.Radius, rl.Black)
	rl.DrawCircle(int32(b.Position.X), int32(b.Position.Y), b.Radius, rl.Black)
}

func (b *Ball) Update(deltaTime float32) {
	b.Position.X += b.Direction.X * deltaTime
	b.Position.Y += b.Direction.Y * deltaTime
}
