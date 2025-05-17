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
const Layer  = struct {
    mapSprite: rl.Texture2D,
    tileSrc:   rl.Rectangle,
    tileDest:  rl.Rectangle,
    tileMap:   []i32,
    mapFile:   []const u8
};
var player: Player = undefined;
var rogue: Npc = undefined;
var layer: Layer = undefined;
const playerSpeed:i32 = 3;
 
fn update() void{
   // if (rl.checkCollisionRecs(player.playerRec, rogue.npcRec) == true ){
   //     player.playerdest.x -= player.playerdest.width;
   //     player.playerRec.x -= player.playerdest.width;

  //  }
    if (    player.playerIsMoving == true){

        switch (player.playerDir) {
                0 => {
                    player.playerdest.y += playerSpeed;
                    player.playerRec.y += playerSpeed;
                },
                1 => {
                    player.playerdest.y -= playerSpeed;
                    player.playerRec.y -= playerSpeed;
                },
                2 => {
                    player.playerdest.x += playerSpeed;
                    player.playerRec.x += playerSpeed;
                },
                3 => {
                    player.playerdest.x -= playerSpeed;
                    player.playerRec.x -= playerSpeed;
                },
        else => {},
        }
    }
    player.playerIsMoving = false;

}
fn input() void {
        if (rl.isKeyDown(.up)) {
            player.playerIsMoving = true;
            player.playerDir = 1;

        }
        if (rl.isKeyDown(.down)) {
            player.playerIsMoving = true;
            player.playerDir = 0;

        }
        if (rl.isKeyDown(.left)) {
            player.playerIsMoving = true;
            player.playerDir = 3;

        }
        if (rl.isKeyDown(.right)) {
            player.playerIsMoving = true;
            player.playerDir = 2;
        }

 
}

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
    const screenwidth = 800;
    const screenheight = 500;

    rl.initWindow(screenwidth, screenheight, "raylib-zig [core] example - basic window");
    defer rl.closeWindow(); // Close window and OpenGL context

    rl.setTargetFPS(60);
    //rl.TraceLog()
    player = Player {
     .playerSprite = try  rl.loadTexture("resources/BasicCharakter.png"),
     .playersrc =  rl.Rectangle{.x = 0, .y = 0, .width = 48, .height = 48},
     .playerdest = rl.Rectangle{.x = 200, .y = 350, .width = 100, .height = 100},
     .playerRec = rl.Rectangle{.x = 135, .y = 280,  .width = 30, .height = 30},
     .playerDir = 4,
     .playerFramecnt = 0,
     .playerIsMoving = true
    };
    rogue = Npc {
     .npcSprite =   try rl.loadTexture("resources/rogue.png"),
     .npcsrc = rl.Rectangle{.x = 0, .y = 0, .width = 32, .height = 32},
     .npcdest = rl.Rectangle{.x = 230, .y = 400, .width = 100, .height = 100},
     .npcRec = rl.Rectangle{.x = 160, .y = 320, .width = 40, .height = 70},
     .npcPrompt = ""
    };
   // layer = Layer {
  //   .mapSprite =   try rl.loadTexture("resources/Grass.png"),
 //    .tileSrc = rl.Rectangle{.x = 0, .y = 0, .width = 16, .height = 16},
  //   .tileDest = rl.Rectangle{.x = 0, .y = 100, .width = 16, .height = 16},
  //   .borderPos = rl.Rectangle{.x = 160, .y = 320, .width = 40, .height = 70},
 //   .mapFile =  try
  //  };
    //-------------------------------------------------------------------------------    // Main game loop
    while (!rl.windowShouldClose()) { // Detect window close button or ESC key
        // Update
        //----------------------------------------------------------------------------------
        // TODO: Update your variables here
        //----------------------------------------------------------------------------------
        input();
        update();
        draw();

    }
}
 