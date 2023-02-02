package gone

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	cam rl.Camera2D
}

var camera_lock = &sync.Mutex{}
var camera_instance *Camera

func Get_Camera() *Camera {
	if camera_instance == nil {
		camera_lock.Lock()
		defer camera_lock.Unlock()

		if camera_instance == nil {
			camera_instance = &Camera{
				cam: rl.NewCamera2D(
					rl.NewVector2(0, 0),
					rl.NewVector2(0, 0),
					0,
					1,
				),
			}
		}
	}

	return camera_instance
}

func (c *Camera) Update_Camera(p *Player) {

	bbox := rl.Vector2{X: 0.3, Y: 0.3}

	bboxMin := rl.GetScreenToWorld2D(
		rl.NewVector2(
			(1-bbox.X)*0.5*float32(W_WIDTH),
			(1-bbox.Y)*0.5*float32(W_HEIGHT),
		),
		c.cam,
	)
	bboxMax := rl.GetScreenToWorld2D(
		rl.NewVector2(
			(1+bbox.X)*0.5*float32(W_WIDTH),
			(1+bbox.Y)*0.5*float32(W_HEIGHT),
		),
		c.cam,
	)
	c.cam.Offset = rl.NewVector2(
		(1-bbox.X)*0.5*float32(W_WIDTH),
		(1-bbox.Y)*0.5*float32(W_HEIGHT),
	)

	if p.pos.X < bboxMin.X {
		c.cam.Target.X = p.pos.X
	}
	if p.pos.Y < bboxMin.Y {
		c.cam.Target.Y = p.pos.Y
	}
	if p.pos.X > bboxMax.X {
		c.cam.Target.X = bboxMin.X + (p.pos.X - bboxMax.X)
	}
	if p.pos.Y > bboxMax.Y {
		c.cam.Target.Y = bboxMin.Y + (p.pos.Y - bboxMax.Y)
	}
}
