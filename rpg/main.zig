/// raylib-zig (c) Nikolas Wipper 2023
const std = @import("std");
const rl = @import("raylib");
const WHITE = rl.Color{ .r = 255, .g = 255, .b = 255, .a = 255 };

const Player = struct {
    playerSprite:  rl.Texture2D,
    playersrc:      rl.Rectangle,
    playerdest:     rl.Rectangle,
    playerRec:      rl.Rectangle,
    playerIsMoving: bool,
    playerDir:      i32,
    playerFramecnt: i32,
};
const  Npc = struct {
    npcSprite: rl.Texture2D,
    npcsrc:    rl.Rectangle,
    npcdest:   rl.Rectangle,
    npcPrompt: []const u8,
    npcRec:    rl.Rectangle
};
var player: Player = undefined;
var rogue: Npc = undefined;
 
 
///pub fn update() void{
///
///}
 

 fn draw() void{
        // Draw
        //----------------------------------------------------------------------------------
        rl.beginDrawing();
        defer rl.endDrawing();

        rl.clearBackground(.white);

        rl.drawText("Congrats! You created your first window!", 190, 200, 20, .light_gray);
        //----------------------------------------------------------------------------------
        rl.drawTexturePro(rogue.npcSprite, rogue.npcsrc, rogue.npcdest, rl.Vector2.init(rogue.npcdest.width, rogue.npcdest.height), 1, WHITE);
        //rl.DrawTexture(grassSprite,100,100,rl.White)
        rl.drawTexturePro(player.playerSprite, player.playersrc, player.playerdest, rl.Vector2.init(player.playerdest.width, player.playerdest.height), 1, WHITE);

}

pub fn main() anyerror!void {
    // Initialization
    //--------------------------------------------------------------------------------------
    const screenWidth = 800;
    const screenHeight = 450;

    rl.initWindow(screenWidth, screenHeight, "raylib-zig [core] example - basic window");
    defer rl.closeWindow(); // Close window and OpenGL context

    rl.setTargetFPS(60);
    //rl.TraceLog()
    player = Player {
     .playerSprite = try  rl.loadTexture("home/fabian/Documents/GO/SproutLands/Basicpack/Characters/BasicCharakter.png"),
     .playersrc =  rl.Rectangle{.x = 0, .y = 0, .width = 48, .height = 48},
     .playerdest = rl.Rectangle{.x = 200, .y = 350, .width = 100, .height = 100},
     .playerRec = rl.Rectangle{.x = 135, .y = 280,  .width = 30, .height = 30},
     .playerDir = 0,
     .playerFramecnt = 0,
     .playerIsMoving = true
    };
    rogue = Npc {
     .npcSprite =   try rl.loadTexture("/home/fabian/Documents/GO/SproutLands/Basicpack/Characters/rogue.png"),
     .npcsrc = rl.Rectangle{.x = 0, .y = 0, .width = 32, .height = 32},
     .npcdest = rl.Rectangle{.x = 230, .y = 400, .width = 100, .height = 100},
     .npcRec = rl.Rectangle{.x = 160, .y = 320, .width = 40, .height = 70},
     .npcPrompt = ""
    };
    //-------------------------------------------------------------------------------    // Main game loop
    while (!rl.windowShouldClose()) { // Detect window close button or ESC key
        // Update
        //----------------------------------------------------------------------------------
        // TODO: Update your variables here
        //----------------------------------------------------------------------------------
        draw();

    }
}

