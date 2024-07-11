package models

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Position     rl.Vector2
	Radius       float32
	Speed        float32
	Acceletation float32
	Direction    rl.Vector2
}

func NewBall(Window rl.Rectangle) *Ball {
	shuffle := [2]float32{-1, 1}

	directionX := shuffle[rand.Int32N(2)] * rand.Float32() * Window.Width
	directionY := shuffle[rand.Int32N(2)] * rand.Float32() * Window.Height

	InitialDirection := rl.NewVector2(directionX, directionY)
	InitialPosition := rl.NewVector2(Window.Width/2, Window.Height/2)

	return &Ball{
		Position:     InitialPosition,
		Radius:       20,
		Speed:        0.5,
		Acceletation: 1.5,
		Direction:    InitialDirection,
	}
}

func (b *Ball) Reset(Window rl.Rectangle) {
	shuffle := [2]float32{-1, 1}
	b.Speed = 0.5
	b.Position = rl.NewVector2(Window.Width/2, Window.Height/2)
	b.Direction = rl.NewVector2(shuffle[rand.Int32N(2)]*rand.Float32()*Window.Width, shuffle[rand.Int32N(2)]*rand.Float32()*Window.Height)
}

func (b *Ball) Draw() {
	rl.DrawCircleV(b.Position, b.Radius, rl.Black)
	rl.DrawCircle(int32(b.Position.X), int32(b.Position.Y), b.Radius, rl.Black)
}

func (b *Ball) Accelerate(deltaTime float32) {
	if b.Speed < 3 {
		b.Speed *= b.Acceletation
	}
}

func (b *Ball) Update(deltaTime float32) {
	b.Position.X += b.Direction.X * b.Speed * deltaTime
	b.Position.Y += b.Direction.Y * b.Speed * deltaTime
}
