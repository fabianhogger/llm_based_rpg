package main

 
import (
  rl "github.com/gen2brain/raylib-go/raylib"
)
const (
	screenWidth = 480
	screenheight = 500
	playerSpeed = 3
)
var (
	running = true
	BkgrColor = rl.NewColor(147,211,196,255)
	grassSprite rl.Texture2D 
	playerSprite rl.Texture2D
	playersrc rl.Rectangle
	playerdest rl.Rectangle
)

func drawScene(){
	rl.DrawTexture(grassSprite,100,100,rl.White)
	rl.DrawTexturePro(playerSprite,playersrc,playerdest, rl.NewVector2(playerdest.Width,	playerdest.Height),1,rl.White)

}
func input(){
	if (rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)){
		playerdest.Y-=playerSpeed
	}
	if (rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown)){
		playerdest.Y+=playerSpeed
	}
	if (rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) ){
		playerdest.X+=playerSpeed
	}
	if (rl.IsKeyDown(rl.KeyL) || rl.IsKeyDown(rl.KeyLeft)){
		playerdest.X-=playerSpeed
	}
	if (rl.IsKeyDown(rl.KeyQ)){
		quit()
	}
}
func update(){
	running= !rl.WindowShouldClose()
}
func render(){		
	rl.BeginDrawing()
	rl.ClearBackground(BkgrColor)
	drawScene()
	rl.EndDrawing()
}
func init(){	
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
  grassSprite = rl.LoadTexture("/home/fabian/Documents/GO/SproutLands/SproutLands _ Sprites _ Basicpack/Tilesets/Grass.png")
  playerSprite = rl.LoadTexture("/home/fabian/Documents/GO/SproutLands/SproutLands _ Sprites _ Basicpack/Characters/Basic Charakter Spritesheet.png")
  playersrc =  rl.NewRectangle(0,0,48, 48)
  playerdest= rl.NewRectangle(200,200,100,100)
}	
func quit(){
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.CloseWindow()
}

func main() {

	for running {
		input()
    update()
    render()
 }
}

