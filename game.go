package main

 
import (
	  rl "github.com/gen2brain/raylib-go/raylib"
	  "fmt"
	  "io/ioutil"
	  "os"
	  "strings"
	  "strconv"
)
const (
	screenWidth = 480
	screenheight = 500
	playerSpeed = 3
)
var (
	running = true
	BkgrColor = rl.NewColor(147,211,196,255)
	grassSprite  rl.Texture2D 
	playerSprite rl.Texture2D
	playersrc    rl.Rectangle
	playerdest   rl.Rectangle
	tileSrc      rl.Rectangle 
	tileDest     rl.Rectangle
	tileMap []int 
	srcMap []string
	borderlist=[8]int{0,1,2,11,13,22,23,24}
	Borderpos []  rl.Rectangle  
	mapW,mapH int
	music rl.Music
	musicPaused bool
	playerFramecnt=0
	Framecount=0
	playerIsMoving= false
	playerDir=0
	val=1
	mapFile= "/home/fabian/Documents/GO/SproutLands/map/grassmap2.csv"

)
func contains(borders [8]int, element int) bool {
    for _, item := range borders {
        if item == element {
            return true
        }
    }
    return false
}

func drawScene(){
	for i:=1; i<len(tileMap); i++{
			tileDest.X = tileDest.Width * float32(i %30)
			tileDest.Y  = tileDest.Height * float32(int(i)/int(30))
	    tileSrc.X = tileSrc.Width * float32(tileMap[i]) 
			tileSrc.Y = tileSrc.Height *float32((tileMap[i])/int(grassSprite.Width/int32(tileSrc.Width)))
			if (contains(borderlist,tileMap[i])==true){
				border:=rl.NewRectangle(tileDest.X,tileDest.Y,8,8)
				Borderpos= append(Borderpos ,border)
 		  	rl.DrawRectangle(int32(border.X),int32(border.Y),int32(border.Width),int32(border.Height),rl.Red)
  
			}
			rl.DrawTexturePro(grassSprite,tileSrc,tileDest, rl.NewVector2(tileDest.Width,	tileDest.Height),1,rl.White)
      
  	 }
 
 	//rl.DrawTexture(grassSprite,100,100,rl.White)
	rl.DrawTexturePro(playerSprite,playersrc,playerdest, rl.NewVector2(playerdest.Width,	playerdest.Height),1,rl.Green)
	rl.DrawRectangle(int32(playerdest.X),int32(playerdest.Y),int32(playerdest.Width),int32(playerdest.Height),rl.Green)

	/*debugg*/
	debug_text:=fmt.Sprintf("Framecount, %d,playerframe %d, entered if %d",Framecount,playerFramecnt,grassSprite.Width)
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
	//collision
	for i:=1;i<len(Borderpos);i++{
  if(rl.CheckCollisionRecs(playerdest,Borderpos[i])){
  	switch(playerDir){
  	case 0:
		 playerdest.Y-=3
    case 1:
		 playerdest.Y+=3	
    case 2:
	  	playerdest.X+=3	
    case 3:
		playerdest.X-=3	
  	break
  }
 }
}

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
func loadMap(mapFile string){
	file,err:= ioutil.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	remNewLines := strings.Replace(string(file),"\n",",",-1)
	sliced := strings.Split(remNewLines,",")
	fmt.Println(sliced)
	fmt.Println(len(sliced))
	mapW =-1
	mapH =-1

	for i:=0;i<len(sliced);i++{
		s,_:= strconv.ParseInt(sliced[i],10,64)
 
		m := int(s)
 
			tileMap = append(tileMap, m)

 
	}
 
}

func init(){	
	rl.InitWindow(1800, 1450, "raylib [core] example - basic window")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
  grassSprite = rl.LoadTexture("/home/fabian/Documents/GO/SproutLands/SproutLands _ Sprites _ Basicpack/Tilesets/Grass.png")
  playerSprite = rl.LoadTexture("/home/fabian/Documents/GO/SproutLands/SproutLands _ Sprites _ Basicpack/Characters/Basic Charakter Spritesheet.png")
  playersrc  =  rl.NewRectangle(0,0,48, 48)
  playerdest = rl.NewRectangle(200,350,100,100)
  rl.InitAudioDevice()
  music=rl.LoadMusicStream("/home/fabian/Documents/GO/SproutLands/SproutLands _ Sprites _ Basicpack/Our-Mountain_v003.mp3")
  rl.PlayMusicStream(music)
  musicPaused= true
  tileDest = rl.NewRectangle(0,0,16,16)
  tileSrc = rl.NewRectangle(0,0,16,16)
 	loadMap(mapFile)

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

