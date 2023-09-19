package gone

import (
	"fmt"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Player_Inputs map[string]int = map[string]int{
	"LEFT":   rl.KeyA,
	"RIGHT":  rl.KeyD,
	"ACTION": rl.KeySpace,
	"OPEN":   rl.KeyW,
}

const (
	Keyboard   = "KY"
	Controller = "CN"
)

type Inputs struct {
	input_type string
}

var inps []int = []int{
	rl.KeyW,
	rl.KeyA,
	rl.KeyK,
	rl.KeyE,
	rl.KeyR,
}

var o_inps []int32 = []int32{}

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

func (i *Inputs) Check_For_Controller() {
	// Check if controller is pluged in

	// Default Keyboard
	if i.input_type != Keyboard {

	}
}

func (i *Inputs) Set_player_Inputs(p *Player) {
	if !rl.IsGamepadAvailable(0) {
		p.inp = Player_Inputs
	} // else set gamepad controller
}

func (i *Inputs) Allow_Player(p *Player) {
	w := 0

	if f := rl.GetKeyPressed(); f != 0 {
		o_inps = append(o_inps, f)
	}

	if !p.Allow_KEY {
		if len(o_inps) == len(inps) {

			for x, i := range o_inps {
				j := int32(inps[x])
				if i == j {

					w++
					if len(inps)-1 == x {
						p.Allow_KEY = true
						fmt.Println("Allow")
					}

					continue
				} else {
					o_inps = []int32{}
					break
				}
			}
		}
	}
}
