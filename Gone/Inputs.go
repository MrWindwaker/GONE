package gone

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Player_Inputs map[string]int = make(map[string]int)

type Inputs struct{}

var input_lock = &sync.Mutex{}
var input_instance *Inputs

func Get_Inputs() *Inputs {
	if input_instance == nil {
		input_lock.Lock()
		defer input_lock.Unlock()

		if input_instance == nil {
			input_instance = &Inputs{}
		}
	}

	return input_instance
}

func (i *Inputs) Set_player_Inputs(p *Player) {
	// Check if controller is pluged in

	// keyboard
	Player_Inputs["LEFT"] = rl.KeyA
	Player_Inputs["RIGHT"] = rl.KeyD

	p.inp = Player_Inputs
}
