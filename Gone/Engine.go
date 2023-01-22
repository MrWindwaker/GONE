package gone

import (
	"fmt"
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
	rl.FlagWindowResizable,
}

type Engine struct {
	Should_Close bool

	// Managers
	inp *Inputs

	player    *Player
	anim_time float32

	coke objs.Trigger
}

var engine_lock = &sync.Mutex{}
var engine_instance *Engine

func Get_Engine() *Engine {
	if engine_instance == nil {
		engine_lock.Lock()
		defer engine_lock.Unlock()

		if engine_instance == nil {
			engine_instance = &Engine{
				Should_Close: false,
				player:       Get_Player(),
				inp:          Get_Inputs(),
				coke: objs.New_Trigger(
					30,
					400,
					rl.NewVector2(500, 500),
					func() {
						fmt.Println("Click")
					},
				),
			}
		}
	}

	return engine_instance
}

func (e *Engine) Init() {
	for _, f := range W_FLAGS {
		rl.SetConfigFlags(f)
	}

	rl.InitWindow(W_WIDTH, W_HEIGHT, W_TITLE)

	e.player.init()
	e.inp.Set_player_Inputs(e.player)
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

	e.coke.Update(e.player)

}

func (e *Engine) Render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.GetColor(0x333333FF))

	e.coke.Render()

	e.player.Render()

	rl.EndDrawing()
}

func (e *Engine) Close() {
	e.player.Close()

	rl.CloseWindow()
}
