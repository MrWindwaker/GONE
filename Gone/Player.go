package gone

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const PLAYER_SPEED int32 = 230

var anims map[string]string = map[string]string{
	"IDLE":  "Assets/Player/balancing.png",
	"MOVE":  "Assets/Player/skip.png",
	"DANCE": "Assets/Player/snap.png",
}

type Player struct {
	pos       rl.Vector2
	is_moving bool
	In_Action bool

	inp map[string]int

	size   int
	width  int
	height int

	animations map[string]rl.Texture2D
	sprite     rl.Texture2D
	c_frame    int
	t_frame    int
	direction  int
}

var player_lock = &sync.Mutex{}
var player_instance *Player

func Get_Player() *Player {
	if player_instance == nil {
		player_lock.Lock()
		defer player_lock.Unlock()

		if player_instance == nil {
			player_instance = &Player{
				pos:        rl.NewVector2(300, 500),
				is_moving:  false,
				size:       2,
				inp:        make(map[string]int),
				animations: make(map[string]rl.Texture2D),
				t_frame:    8,
				c_frame:    0,
				direction:  -1,
			}
		}
	}

	return player_instance
}

func (p *Player) init() {
	for name, path := range anims {
		p.animations[name] = rl.LoadTexture(path)
	}

	p.sprite = p.animations["IDLE"]
	p.width = int(p.sprite.Width) / p.t_frame
	p.height = int(p.sprite.Height)
}

func (p *Player) move(dt float32) {
	if !p.In_Action {
		if rl.IsKeyDown(int32(p.inp["RIGHT"])) {
			p.is_moving = true
			p.pos.X += float32(PLAYER_SPEED) * dt

			p.change_animation("MOVE")
			p.change_direction(-1)
		}

		if rl.IsKeyDown(int32(p.inp["LEFT"])) {
			p.is_moving = true
			p.pos.X -= float32(PLAYER_SPEED) * dt

			p.change_animation("MOVE")
			p.change_direction(1)
		}

		if rl.IsKeyDown(rl.KeySpace) {
			p.change_animation("DANCE")
			p.In_Action = true
		}
	}

	if rl.IsKeyUp(rl.KeySpace) && p.In_Action {
		p.In_Action = false
	}

	if !p.is_moving && !p.In_Action {
		p.change_animation("IDLE")
	}

	if rl.IsKeyUp(int32(p.inp["RIGHT"] | p.inp["LEFT"])) {
		p.is_moving = false
	}
}

func (p *Player) get_source() rl.Rectangle {
	return rl.NewRectangle(
		float32(p.c_frame)*float32(p.width),
		0,
		float32(p.width)*float32(p.direction),
		float32(p.height),
	)
}

func (p *Player) get_dest() rl.Rectangle {
	return rl.NewRectangle(
		p.pos.X,
		p.pos.Y,
		float32(p.width)*float32(p.size),
		float32(p.height)*float32(p.size),
	)
}

func (p *Player) animate() {
	p.c_frame++

	if p.c_frame >= p.t_frame {
		p.c_frame = 0
	}
}

func (p *Player) change_animation(name string) {
	p.sprite = p.animations[name]
}

func (p *Player) change_direction(dir int) {
	p.direction = dir
}

func (p *Player) Update(dt float32) {
	p.move(dt)
}

func (p *Player) Render() {
	rl.DrawTexturePro(
		p.sprite,
		p.get_source(),
		p.get_dest(),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}

func (p *Player) Close() {
	for _, t := range p.animations {
		rl.UnloadTexture(t)
	}
}

func (p *Player) Get_Rec() rl.Rectangle {
	return p.get_dest()
}

func (p *Player) Doing_Action() {
	p.In_Action = true
}

func (p *Player) Finish_Action() {
	p.In_Action = false
}
