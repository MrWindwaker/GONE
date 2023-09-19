package gone

import (
	"sync"

	objs "wasm/game/Gone/Objects"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const W_WIDTH int32 = 1280
const W_HEIGHT int32 = 720
const W_TITLE string = "GONE"

const ANIMATION_SPEED float32 = 1.0 / 12.0

var W_FLAGS []uint32 = []uint32{
	rl.FlagMsaa4xHint,
	rl.FlagVsyncHint,
}

type Engine struct {
	Should_Close bool
	Settings     Settings

	// Managers
	inp *Inputs
	cm  *Camera
	ui  *GUI
	sm  *SceneManager

	player    *Player
	anim_time float32

	fls    map[string]objs.Floor
	curr_f string
}

var engine_lock = &sync.Mutex{}
var engine_instance *Engine

func is_dev() bool {
	return engine_instance.Settings.Is_Dev
}

func Get_Engine() *Engine {
	if engine_instance == nil {
		engine_lock.Lock()
		defer engine_lock.Unlock()

		if engine_instance == nil {

			engine_instance = &Engine{
				Should_Close: false,
				player:       Get_Player(),
				inp:          Get_Inputs(),
				cm:           Get_Camera(),
				ui:           Get_GUI(),
				Settings:     Load_Settings(),
				sm:           get_scene_manager(),
				fls:          make(map[string]objs.Floor),
				curr_f:       "F1",
			}
		}
	}

	return engine_instance
}

func (e *Engine) Init() {

	Check_Wall()
	e.Settings = Load_Settings()

	for _, f := range W_FLAGS {
		rl.SetConfigFlags(f)
	}

	rl.InitWindow(W_WIDTH, W_HEIGHT, W_TITLE)
	if !e.Settings.Is_Dev {
		rl.SetExitKey(0)
	}

	e.player.init()
	e.inp.Set_player_Inputs(e.player)

	e.sm.Init()
}

func (e *Engine) Update() {
	e.Should_Close = rl.WindowShouldClose()
	dt := rl.GetFrameTime()

	e.anim_time += dt

	if e.anim_time >= ANIMATION_SPEED {
		e.anim_time = 0
		e.player.animate()
	}

	e.player.Update(dt)

	e.cm.Update_Camera(e.player)

	e.inp.Allow_Player(e.player)
}

func (e *Engine) Render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.GetColor(0x333333FF))

	rl.BeginMode2D(e.cm.cam)

	rl.DrawRectangle(
		W_WIDTH+500,
		500,
		100,
		150,
		rl.Blue,
	)

	rl.DrawText(
		"HTTPS://LTTSTORE.COM",
		10,
		1500,
		20,
		rl.GetColor(0xC73A03FF),
	)

	// Draw floor and background here
	rl.DrawRectangle(500, 500, 1200, 300, rl.Gold)

	e.sm.Render(e.player)

	// Draw foreground Elements here
	rl.DrawRectangle(500, 600, 1200, 30, rl.Blue)

	rl.EndMode2D()

	e.ui.Render()

	rl.EndDrawing()
}

func (e *Engine) Close() {
	e.player.Close()

	rl.CloseWindow()
}
