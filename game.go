package main

 
import (
	  rl "github.com/gen2brain/raylib-go/raylib"
	  "fmt"
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
	music rl.Music
	musicPaused bool
	playerFramecnt=0
	Framecount=0
	playerIsMoving= false
	playerDir=0
	val=1
)

func drawScene(){
	rl.DrawTexture(grassSprite,100,100,rl.White)
	rl.DrawTexturePro(playerSprite,playersrc,playerdest, rl.NewVector2(playerdest.Width,	playerdest.Height),1,rl.White)
	/*debugg*/
	debug_text:=fmt.Sprintf("Framecount, %d,playerframe %d, entered if %d",Framecount,playerFramecnt,val)
	rl.DrawText(debug_text,150,50,10,rl.White)
  
}
func input(){


	if (rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp)){
		playerIsMoving = true
		playerDir=1
		
	}
	if (rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown)){
		playerIsMoving = true
		playerDir=0
		
	}
	if (rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) ){
		playerIsMoving = true
		playerDir=3
		
	}
	if (rl.IsKeyDown(rl.KeyL) || rl.IsKeyDown(rl.KeyLeft)){
		playerIsMoving = true
		playerDir=2
	

	}
	if (rl.IsKeyDown(rl.KeyQ)){
		quit()
	}
	if(rl.IsKeyPressed(rl.KeyP)){
		musicPaused = !musicPaused
	}
}
func update(){
	rl.UpdateMusicStream(music)
	rl.ResumeMusicStream(music)
	running= !rl.WindowShouldClose()
	playersrc.X=playersrc.Width*float32(playerFramecnt)

	Framecount++
	if Framecount==65 {Framecount=0}
	if playerIsMoving{
	switch(playerDir){
	case 0:
		playerdest.Y+=playerSpeed	
  case 1:
		playerdest.Y-=playerSpeed	
  case 2:
		playerdest.X-=playerSpeed	
  case 3:
		playerdest.X+=playerSpeed	
	}
	if  Framecount%8==1 { 
		playerFramecnt++
}  


}
//idle animation
	if !playerIsMoving && Framecount%45==1{
		if 	playerFramecnt==1{
			playerFramecnt=0
		}else{
			playerFramecnt=1
		} 

}

	
	if playerFramecnt>3  {
		playerFramecnt=0
	}
	playersrc.X=playersrc.Width*float32(playerFramecnt)
	playersrc.Y=playersrc.Height*float32(playerDir)
	playerIsMoving=false

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
  playersrc  =  rl.NewRectangle(0,0,48, 48)
  playerdest = rl.NewRectangle(200,200,100,100)
  rl.InitAudioDevice()
  music=rl.LoadMusicStream("/home/fabian/Documents/GO/SproutLands/SproutLands _ Sprites _ Basicpack/Our-Mountain_v003.mp3")
  rl.PlayMusicStream(music)
  musicPaused= false
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

