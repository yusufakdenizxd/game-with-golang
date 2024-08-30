package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running      = true
	bkgColor     = rl.NewColor(147, 211, 196, 255)
	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D
	playerSrc    rl.Rectangle
	playerDest   rl.Rectangle
	playerSpeed  float32 = 3
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) {
		playerDest.Y -= playerSpeed
	}

	if rl.IsKeyDown(rl.KeyS) {
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) {
		playerDest.X += playerSpeed
	}

	if rl.IsKeyDown(rl.KeyA) {
		playerDest.X -= playerSpeed
	}
}
func update() {
	running = !rl.WindowShouldClose()
}
func render() {
	rl.BeginDrawing()
	rl.ClearBackground(bkgColor)

	drawScene()

	rl.EndDrawing()
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.CloseWindow()
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "SproutSling")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	grassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
	playerSprite = rl.LoadTexture("res/Characters/basicCharacterSpritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	for running {
		input()
		update()
		render()
	}
}
