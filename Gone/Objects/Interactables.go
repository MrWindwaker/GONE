package objects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Interactable struct {
	Active     bool
	active_key int

	width  int
	height int
	pos    rl.Vector2
	color  uint

	script fn
}

func New_Interactable(pos rl.Vector2, width int, height int, key int, sc fn, color uint) Interactable {
	return Interactable{
		Active:     false,
		active_key: key,
		script:     sc,
		width:      width,
		height:     height,
		pos:        pos,
		color:      color,
	}
}

func (i *Interactable) run() {
	if i.Active {
		i.script()
	}
}

func (i *Interactable) Get_Rec() rl.Rectangle {
	return rl.NewRectangle(
		i.pos.X,
		i.pos.Y,
		float32(i.width),
		float32(i.height),
	)
}

func (i *Interactable) Update(p pl) {
	if rl.CheckCollisionRecs(i.Get_Rec(), p.Get_Collision()) {
		i.Active = true
	} else {
		i.Active = false
	}

	if i.Active && rl.IsKeyPressed(int32(i.active_key)) {
		i.run()
	}
}

func (i *Interactable) Render() {
	rl.DrawRectangleRec(
		i.Get_Rec(),
		rl.GetColor(i.color),
	)
}
