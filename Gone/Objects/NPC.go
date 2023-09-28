package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type NPC struct {
	Pos    rl.Vector2
	Sprite rl.Texture2D

	has_anim bool
}

func Create_NPC(fl string, x, y float32) NPC {
	return NPC{
		Pos:    rl.NewVector2(0, 0),
		Sprite: rl.LoadTexture(fl),
	}
}

func (n *NPC) Unload() {
	rl.UnloadTexture(n.Sprite)
}
