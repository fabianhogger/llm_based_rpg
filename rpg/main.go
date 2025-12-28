package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"io/ioutil"
	"log"
	"os"
	ml "rpg/prompt_model"
	"strconv"
	"strings"
	"time"
)

const (
	screenWidth  = 480
	screenheight = 500
	playerSpeed  = 3
)

type Player struct {
	playerSprite   rl.Texture2D
	playersrc      rl.Rectangle
	playerdest     rl.Rectangle
	playerRec      rl.Rectangle
	playerIsMoving bool
	playerDir      int
	playerFramecnt int
}
type Npc struct {
	npcSprite rl.Texture2D
	npcsrc    rl.Rectangle
	npcdest   rl.Rectangle
	npcPrompt string
	npcRec    rl.Rectangle
}

type Layer struct {
	mapSprite rl.Texture2D
	tileSrc   rl.Rectangle
	tileDest  rl.Rectangle
	tileMap   []int
	srcMap    []string
	Borderpos []rl.Rectangle
	mapFile   string
}

var (
	running             = true
	BkgrColor           = rl.NewColor(147, 211, 196, 255)
	mapborder      bool = true
	borderlist          = [8]int{0, 1, 2, 11, 13, 22, 23, 24}
	mapW, mapH     int
	music          rl.Music
	musicPaused    bool
	colarea        rl.Rectangle
	Framecount     = 0
	mapFile        = "../map/grassmap2.csv"
	player         Player
	layer          Layer
	rogue          Npc
	textInput      string
	writing        bool
	dialog         bool
	dialogActive   bool
	dialogCooldown int
	replychain     string
	loggerFileName = "./logger.text"
	chatborder     = 0
)

func contains(borders [8]int, element int) bool {
	for _, item := range borders {
		if item == element {
			return true
		}
	}
	return false
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "unknown"
	}
	return dir
}

func npcReply(character Npc, question string) string {
	answer := ml.Ask(question)
	log.Println(answer)
	return answer
}
func readInput() {
	time.Sleep(100 * time.Millisecond)
	writing = true
	for !rl.IsKeyDown(rl.KeyEnter) {
		// Handle backspace
		if rl.IsKeyPressed(rl.KeyBackspace) && len(textInput) > 0 {
			textInput = textInput[:len(textInput)-1]
		}

		// Get character input
		char := rl.GetCharPressed()
		if char >= 32 && char <= 126 {
			textInput += string(char)
		}
	}

	if textInput != "" {
		log.Println(textInput)
		chatborder += 100
		reply := npcReply(rogue, "I ask you "+textInput+" reply like a small town farmer in rpg game with a short phrase")
		replychain += ">" + textInput + "\n" + "Farmer:" + reply + "\n"
	}

	writing = false
	dialogActive = false
	textInput = ""
}

func drawLayer() {
	var borderArr []rl.Rectangle

	for i := 0; i < len(layer.tileMap); i++ {
		layer.tileDest.X = layer.tileDest.Width * float32(i%30)
		layer.tileDest.Y = layer.tileDest.Height * float32(int(i)/int(30))
		layer.tileSrc.X = layer.tileSrc.Width * float32(layer.tileMap[i])
		layer.tileSrc.Y = layer.tileSrc.Height * float32((layer.tileMap[i])/int(layer.mapSprite.Width/int32(layer.tileSrc.Width)))

		if contains(borderlist, layer.tileMap[i]) {
			if mapborder {
				border := rl.NewRectangle(layer.tileDest.X, layer.tileDest.Y, layer.tileDest.Width, layer.tileDest.Height)
				borderArr = append(borderArr, border)
			}
		}
		rl.DrawTexturePro(layer.mapSprite, layer.tileSrc, layer.tileDest, rl.NewVector2(0, 0), 0, rl.White)
	}

	if mapborder {
		layer.Borderpos = borderArr
	}
}
func drawScene() {
	drawLayer()

	mapborder = false
	rl.DrawTexturePro(rogue.npcSprite, rogue.npcsrc, rogue.npcdest, rl.NewVector2(rogue.npcdest.Width, rogue.npcdest.Height), 1, rl.White)
	//rl.DrawTexture(grassSprite,100,100,rl.White)
	rl.DrawTexturePro(player.playerSprite, player.playersrc, player.playerdest, rl.NewVector2(player.playerdest.Width, player.playerdest.Height), 1, rl.White)
	//rl.DrawRectangle(int32(player.playerRec.X ),int32(player.playerRec.Y),int32(player.playerRec.Width),int32(player.playerRec.Height),rl.Red)
	//rl.DrawRectangle(int32(rogue.npcRec.X),int32(rogue.npcRec.Y),int32(rogue.npcRec.Width),int32(rogue.npcRec.Height),rl.Blue)
	if dialog {
		rl.DrawRectangle(150, 450, int32(200+chatborder), int32(50+chatborder), rl.Black)
		if writing {
			rl.DrawText(">"+textInput, 150, int32(rl.GetScreenHeight()-1000), 50, rl.White)
		} else {
			rl.DrawText(replychain, 150, int32(rl.GetScreenHeight()-1000), 50, rl.White)
			rl.DrawText("Press R to talk, Q to exit", 150, 420, 20, rl.Yellow)
		}
	}

	/*debugg*/
	debug_text := fmt.Sprintf("Framecount, %d,playerframe %d, entered if %d", Framecount, player.playerFramecnt, layer.mapSprite.Width)
	rl.DrawText(debug_text, 150, 50, 10, rl.White)
}

func input() {
	if dialog {
		if rl.IsKeyPressed(rl.KeyR) && !writing && !dialogActive {
			go readInput()
		} else if rl.IsKeyPressed(rl.KeyQ) && !writing {
			replychain = ""
			dialog = false
			dialogActive = false
			dialogCooldown = 30
			chatborder = 0
		}
	} else {
		if rl.IsKeyDown(rl.KeyUp) {
			player.playerIsMoving = true
			player.playerDir = 1

		}
		if rl.IsKeyDown(rl.KeyDown) {
			player.playerIsMoving = true
			player.playerDir = 0

		}
		if rl.IsKeyDown(rl.KeyRight) {
			player.playerIsMoving = true
			player.playerDir = 3

		}
		if rl.IsKeyDown(rl.KeyLeft) {
			player.playerIsMoving = true
			player.playerDir = 2
		}
		if rl.IsKeyPressed(rl.KeyP) {
			musicPaused = !musicPaused
		}
	}
}

func update() {
	rl.UpdateMusicStream(music)
	rl.ResumeMusicStream(music)
	running = !rl.WindowShouldClose()

	// Update dialog cooldown
	if dialogCooldown > 0 {
		dialogCooldown--
	}

	//collision
	collision := false
	if rl.CheckCollisionRecs(player.playerRec, rogue.npcRec) && !dialogActive && dialogCooldown == 0 {
		dialogActive = true
		dialog = true
		log.Println("collision detected - starting dialog")
		go readInput()

		// Push player back based on direction
		switch player.playerDir {
		case 0: // down
			player.playerdest.Y -= player.playerdest.Height
			player.playerRec.Y -= player.playerRec.Height
		case 1: // up
			player.playerdest.Y += player.playerdest.Height
			player.playerRec.Y += player.playerRec.Height
		case 2: // left
			player.playerdest.X += player.playerdest.Width
			player.playerRec.X += player.playerRec.Width
		case 3: // right
			player.playerdest.X -= player.playerdest.Width
			player.playerRec.X -= player.playerdest.Width
		}
	}
	for i := 0; i < len(layer.Borderpos); i++ {
		if rl.CheckCollisionRecs(player.playerRec, layer.Borderpos[i]) {
			switch player.playerDir {
			case 0:
				player.playerdest.Y -= player.playerdest.Height
				player.playerRec.Y -= player.playerRec.Height
			case 1:
				player.playerdest.Y += player.playerdest.Height
				player.playerRec.Y += player.playerRec.Height
			case 2:
				player.playerdest.X += player.playerdest.Width
				player.playerRec.X += player.playerRec.Width
			case 3:
				player.playerdest.X -= player.playerdest.Width
				player.playerRec.X -= player.playerRec.Width
			}
			collision = true
			break
		}
	}

	Framecount++
	if Framecount == 65 {
		Framecount = 0
	}
	if player.playerIsMoving && !collision {
		switch player.playerDir {
		case 0:
			player.playerdest.Y += playerSpeed
			player.playerRec.Y += playerSpeed

		case 1:
			player.playerdest.Y -= playerSpeed
			player.playerRec.Y -= playerSpeed

		case 2:
			player.playerdest.X -= playerSpeed
			player.playerRec.X -= playerSpeed

		case 3:
			player.playerdest.X += playerSpeed
			player.playerRec.X += playerSpeed

		}
		if Framecount%8 == 1 {
			player.playerFramecnt++
		}
	}

	//idle animation
	if !player.playerIsMoving && Framecount%45 == 1 {
		if player.playerFramecnt == 1 {
			player.playerFramecnt = 0
		} else {
			player.playerFramecnt = 1
		}
	}

	if player.playerFramecnt > 3 {
		player.playerFramecnt = 0
	}
	player.playersrc.X = player.playersrc.Width * float32(player.playerFramecnt)
	player.playersrc.Y = player.playersrc.Height * float32(player.playerDir)
	player.playerIsMoving = false

}
func render() {
	rl.BeginDrawing()
	rl.ClearBackground(BkgrColor)
	drawScene()
	rl.EndDrawing()
}
func loadMap(mapFile string) []int {
	file, err := ioutil.ReadFile(mapFile)
	if err != nil {
		log.Printf("ERROR: Failed to load map file '%s': %v\n", mapFile, err)
		log.Printf("Current working directory: %s\n", getCurrentDir())
		log.Fatal(err)
		os.Exit(1)
	}
	remNewLines := strings.Replace(string(file), "\n", ",", -1)
	sliced := strings.Split(remNewLines, ",")
	var tileMaparr []int
	for i := 0; i < len(sliced); i++ {
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		tileMaparr = append(tileMaparr, m)

	}
	return tileMaparr

}

func init() {
	logFile, err := os.OpenFile(loggerFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Unable to create Logger file:", err.Error())
		return
	}
	log.SetOutput(logFile)
	rl.InitWindow(1800, 1450, "raylib [core] example - basic window")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
	player.playerSprite = rl.LoadTexture("./Basicpack/Characters/BasicCharakter.png")
	player.playersrc = rl.NewRectangle(0, 0, 48, 48)
	player.playerdest = rl.NewRectangle(200, 350, 100, 100)
	// Center collision box within sprite (200 + 35, 350 + 35, smaller box)
	player.playerRec = rl.NewRectangle(235, 385, 30, 30)
	//init npc
	rogue.npcSprite = rl.LoadTexture("./Basicpack/Characters/rogue.png")
	rogue.npcsrc = rl.NewRectangle(0, 0, 32, 32)
	rogue.npcdest = rl.NewRectangle(230, 400, 100, 100)
	// Center collision box within sprite (230 + 30, 400 + 30, smaller box)
	rogue.npcRec = rl.NewRectangle(260, 430, 40, 40)
	//music
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("./Basicpack/Our-Mountain_v003.mp3")
	rl.PlayMusicStream(music)
	musicPaused = true
	//music
	//map
	layer.mapSprite = rl.LoadTexture("./Basicpack/Tilesets/Grass.png")
	layer.tileDest = rl.NewRectangle(0, 100, 16, 16)
	layer.tileSrc = rl.NewRectangle(0, 0, 16, 16)
	layer.tileMap = loadMap(mapFile)
	replychain = ""
}
func quit() {
	rl.UnloadTexture(layer.mapSprite)
	rl.UnloadTexture(player.playerSprite)
	rl.UnloadTexture(rogue.npcSprite)
	rl.CloseWindow()
}

func main() {
	for running {
		input()
		update()
		render()
	}
}
