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
