package main

import (
	"fmt"
	m "game/pong/models"

	rl "github.com/gen2brain/raylib-go/raylib"

	s "game/pong/services"
)

func main() {
	screenWidth := 1920
	screenHeight := 1080
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Pong")

	// Set our game to run at 60 frames-per-second
	rl.SetTargetFPS(60)

	// Load logo
	logo := rl.LoadImage("Assets/logo.png")
	rl.SetWindowIcon(*logo)

	// Load background
	bg := rl.LoadTexture("Assets/bg.png")
	//container
	bgSize := rl.NewRectangle(0, 0, float32(bg.Width), float32(bg.Height))
	//position
	bgPos := rl.NewVector2(0, 0)
	//window size
	windowSize := rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight))

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Create ball
	ball := m.NewBall(windowSize)

	player := m.NewPadle(windowSize, true)
	enemy := m.NewPadle(windowSize, false)
	score := m.NewScore()

	for !rl.WindowShouldClose() {
		//m.MoveUp(rl.GetFrameTime())
		ball.Update(rl.GetFrameTime())
		player.Move(rl.GetFrameTime(), windowSize)

		if int32(rl.GetFrameTime())%10 == 0 {
			enemy.Bot(rl.GetFrameTime(), windowSize, ball.Position)
		}

		rl.BeginDrawing()

		s.CollisionService(windowSize, ball, player, enemy, score)
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexturePro(bg, bgSize, windowSize, bgPos, 0.0, rl.White)

		player.Draw()
		enemy.Draw()
		ball.Draw()

		// Draw player score
		rl.DrawText(fmt.Sprint(score.Player), int32(screenWidth/2+25), 20, 40, rl.White)

		// Draw enemy score
		rl.DrawText(fmt.Sprint(score.Enemy), int32(screenWidth/2-50), 20, 40, rl.White)

		rl.EndDrawing()
	}

}
