package gone

import (
	"sync"

	objs "wasm/game/Gone/Objects"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const W_TITLE string = "GONE"

const ANIMATION_SPEED float32 = 1.0 / 12.3
const BG_ANIM float32 = 1.0 / 9.0

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

	// Delete later
	fl map[string]objs.Floor
	np objs.NPC
	bg objs.BackgroundAnimted

	W_WIDTH  int
	W_HEIGHT int
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
	e.Settings = Load_Settings()

	for _, f := range W_FLAGS {
		rl.SetConfigFlags(f)
	}

	rl.InitWindow(20, 20, W_TITLE)

	e.W_WIDTH = rl.GetMonitorWidth(0)
	e.W_HEIGHT = rl.GetMonitorHeight(0)

	rl.SetWindowPosition(0, 0)
	rl.SetWindowSize(int(e.W_WIDTH), int(e.W_HEIGHT))
	rl.SetTargetFPS(60)
	rl.ToggleFullscreen()

	if !is_dev() {
		rl.SetExitKey(0)
	}

	if !e.Settings.Is_Dev {
		rl.SetExitKey(0)
	}

	e.player.init()
	e.inp.Set_player_Inputs(e.player)

	e.sm.Init()

	e.bg = objs.BackgroundAnimted{
		Texture:     rl.LoadTexture("Assets/Backgrounds/city_night/city_full.png"),
		Is_Animated: false,
	}
	e.fl = create_floor()
	e.np = objs.Create_NPC("Assets/NPC/mercy.jpg", 45, 300)
}

func (e *Engine) Update() {
	e.Should_Close = rl.WindowShouldClose()
	dt := rl.GetFrameTime()

	e.anim_time += dt

	if e.anim_time >= ANIMATION_SPEED {

		e.bg.Animate()

		e.anim_time = 0
		e.player.animate()

	}

	e.player.Update(dt)

	e.cm.Update_Camera(e.player, e.W_WIDTH, e.W_HEIGHT)

	e.inp.Allow_Player(e.player)
}

func (e *Engine) Render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.GetColor(0x333333FF))

	e.bg.Render()

	rl.BeginMode2D(e.cm.cam)

	rl.DrawRectangle(
		int32(e.W_WIDTH+500),
		500,
		100,
		150,
		rl.Blue,
	)

	rl.DrawText(
		"Arturo Was Here",
		10,
		1500,
		20,
		rl.GetColor(0xC73A03FF),
	)

	// Draw floor
	for _, f := range e.fl {
		f.Render()
	}

	e.sm.Render(e.player)

	// Draw foreground Elements here
	//rl.DrawTexture(e.np.Sprite, int32(e.np.Pos.X), 550, rl.White)
	rl.EndMode2D()

	e.ui.Render()

	if is_dev() {
		rl.DrawFPS(20, 20)
	}

	rl.EndDrawing()
}

func (e *Engine) Close() {
	e.player.Close()
	e.np.Unload()

	rl.UnloadTexture(e.bg.Texture)

	rl.CloseWindow()
}
