package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type Trigger struct {
	width  int
	height int
	pos    rl.Vector2

	is_active bool
	script    fn
}

func New_Trigger(width int, height int, pos rl.Vector2, sc fn) Trigger {
	return Trigger{
		width:     width,
		height:    height,
		pos:       pos,
		script:    sc,
		is_active: true,
	}
}

func (t *Trigger) Get_Rec() rl.Rectangle {
	return rl.NewRectangle(
		t.pos.X,
		t.pos.Y,
		float32(t.width),
		float32(t.height),
	)
}

func (t *Trigger) Update(p pl) {
	if rl.CheckCollisionRecs(t.Get_Rec(), p.Get_Rec()) && t.is_active {
		t.script()

		t.is_active = false
	}
}

func (t *Trigger) Render() {
	rl.DrawRectangleRec(
		t.Get_Rec(),
		rl.GetColor(0x00000000),
	)
}
