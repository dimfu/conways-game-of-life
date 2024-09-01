package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const (
		WINDOW_HEIGHT int32 = 500
		WINDOW_WIDTH  int32 = 500
		FPS           int32 = 16
	)

	var (
		bgColor rl.Color = rl.Color{R: 60, G: 56, B: 54, A: 255}
	)

	sim := NewSim(WINDOW_HEIGHT, WINDOW_WIDTH, 4)

	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "conway's game of life")
	rl.SetTargetFPS(FPS)

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		sim.Update()
		rl.BeginDrawing()
		rl.ClearBackground(bgColor)
		sim.Draw()
		rl.EndDrawing()
	}
}
