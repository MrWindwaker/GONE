package gone

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GUI struct{}

var gui_lock = &sync.Mutex{}
var gui_instance *GUI

func Get_GUI() *GUI {
	if gui_instance == nil {
		gui_lock.Lock()
		defer gui_lock.Unlock()

		if gui_instance == nil {
			gui_instance = &GUI{}
		}
	}

	return gui_instance
}

func (g *GUI) Render() {
	rl.DrawRectangleRounded(
		rl.NewRectangle(10, 10, 300, 150),
		0.2,
		0,
		rl.Green,
	)
}
