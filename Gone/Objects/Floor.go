package objects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Floor struct {
	pos    rl.Vector2
	width  int
	height int
	offset int

	cam_fix int
}

type FloorsJson struct {
	Floors []FloorJson `json:"floors"`
}

type FloorJson struct {
	Name   string  `json:"name"`
	Pos    PosJson `json:"pos"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Offset int     `json:"offset"`
}

type PosJson struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func New_Floor(pos rl.Vector2, w int, h int, offset int) Floor {
	return Floor{
		pos:     pos,
		width:   w,
		height:  h,
		offset:  offset,
		cam_fix: 100,
	}
}

func (f *Floor) Get_Rec() rl.Rectangle {
	return rl.NewRectangle(
		f.pos.X,
		f.pos.Y,
		float32(f.width),
		float32(f.height)+float32(f.cam_fix),
	)
}

func (f *Floor) Get_Collision() rl.Rectangle {
	return rl.NewRectangle(
		f.pos.X,
		f.pos.Y+float32(f.offset),
		float32(f.width),
		float32(f.height)-float32(f.offset)+float32(f.cam_fix),
	)
}

func (f *Floor) Render() {
	rl.DrawRectangleRec(
		f.Get_Rec(),
		rl.Red,
	)

	rl.DrawRectangleRec(
		f.Get_Collision(),
		rl.GetColor(0xFFD70088),
	)
}
