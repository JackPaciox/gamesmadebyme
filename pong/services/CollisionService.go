package services

import (
	m "game/pong/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// foreach frame check if the ball collides with the window, player, enemy or the goal
func CollisionService(Window rl.Rectangle, b *m.Ball, player *m.Padle, enemy *m.Padle, score *m.Score) {

	//check if the ball collides with something
	switch {
	//check if the ball collides with the top of window
	case rl.CheckCollisionCircleRec(b.Position, b.Radius, rl.NewRectangle(0, Window.Height, Window.Width, 1)):
		b.Direction.Y *= -1
	//check if the ball collides with the bottom of window
	case rl.CheckCollisionCircleRec(b.Position, b.Radius, rl.NewRectangle(0, 0, Window.Width, 1)):
		b.Direction.Y *= -1
	//check if the ball collides with the player
	case rl.CheckCollisionCircleRec(b.Position, b.Radius, rl.NewRectangle(player.Position.X, player.Position.Y, player.Width, player.Height)):
		b.Direction.X *= -1
	//check if the ball collides with the enemy
	case rl.CheckCollisionCircleRec(b.Position, b.Radius, rl.NewRectangle(enemy.Position.X, enemy.Position.Y, enemy.Width, enemy.Height)):
		b.Direction.X *= -1
	//check if the player scored
	case rl.CheckCollisionCircleRec(b.Position, b.Radius, rl.NewRectangle(0, 0, 1, Window.Height)):
		score.PlayerScored()
		b.Reset(Window)
	//check if the enemy scored
	case rl.CheckCollisionCircleRec(b.Position, b.Radius, rl.NewRectangle(Window.Width, 0, 1, Window.Height)):
		score.EnemyScored()
		b.Reset(Window)
	}
}
