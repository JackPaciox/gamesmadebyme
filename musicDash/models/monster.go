package models

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Monster struct {
	Position rl.Vector2
	Width    float32
	Height   float32
	Speed    float32
}

func NewMonster(Window rl.Rectangle) *Monster {
	monsterWidth := float32(480)
	monsterHeight := float32(1080)
	initialPosition := rl.Vector2{}
	return &Monster{
		Position: initialPosition,
		Width:    monsterWidth,
		Height:   monsterHeight,
		Speed:    600,
	}
}

func (m *Monster) Draw() {
	rl.DrawRectangleV(m.Position, rl.NewVector2(m.Width, m.Height), rl.White)
}
