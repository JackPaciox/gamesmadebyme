package models

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Score struct {
	Player int
	Enemy  int
}

func NewScore() *Score {
	return &Score{
		Player: 0,
		Enemy:  0,
	}
}

func (s *Score) PlayerScored() {
	s.Player++
}

func (s *Score) EnemyScored() {
	s.Enemy++
}

func (s *Score) Reset() {
	s.Player = 0
	s.Enemy = 0
}

func Draw(s *Score, screenWidth int32) {
	// Draw player score
	rl.DrawText(fmt.Sprintf("%d", s.Player), screenWidth/2-50, 20, 40, rl.White)
	// Draw enemy score
	rl.DrawText(fmt.Sprintf("%d", s.Enemy), screenWidth/2+25, 20, 40, rl.White)
}
