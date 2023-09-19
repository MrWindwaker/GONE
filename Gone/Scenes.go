package gone

import (
	"sync"
	objects "wasm/game/Gone/Objects"
)

type Scene struct {
	flrs objects.Floor
}

type SceneManager struct {
	scenes map[string]Scene
}

var sm_lock = &sync.Mutex{}
var sm_instance *SceneManager

func get_scene_manager() *SceneManager {

	if sm_instance == nil {
		sm_lock.Lock()
		defer sm_lock.Unlock()

		if sm_instance == nil {
			sm_instance = &SceneManager{}
		}
	}

	return sm_instance
}

func (sm *SceneManager) Init() {

}

func (sm *SceneManager) Render(pl *Player) {
	//Render Scene

	pl.Render()
}
