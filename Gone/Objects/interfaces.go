package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type fn func()

type Object interface {
	Render()
	Update()
	//Get_Rec() rl.Rectangle
}

type pl interface {
	Get_Rec() rl.Rectangle
	Doing_Action()
	Finish_Action()
}
