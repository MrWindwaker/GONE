package objects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Floor struct {
	Pos    rl.Vector2
	Width  int
	Height int
	Offset int
	Name   string
	Color  int

	cam_fix int
}

func Get_A_Floor() []Floor {
	var fl []Floor

	return fl
}

type FJSON struct {
	Floors []FloorJson `json:"floors"`
}

type FloorJson struct {
	Name   string  `json:"name"`
	Pos    PosJson `json:"pos"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Offset int     `json:"offset"`
	Color  string  `json:"color"`
}

type PosJson struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func New_Floor(pos rl.Vector2, w int, h int, offset int, name string, color int) Floor {
	return Floor{
		Pos:     pos,
		Width:   w,
		Height:  h,
		Offset:  offset,
		cam_fix: 100,
		Name:    name,
		Color:   color,
	}
}

func (f *Floor) Get_Rec() rl.Rectangle {
	return rl.NewRectangle(
		f.Pos.X,
		f.Pos.Y,
		float32(f.Width),
		float32(f.Height)+float32(f.cam_fix),
	)
}

func (f *Floor) Get_Collision() rl.Rectangle {
	return rl.NewRectangle(
		f.Pos.X,
		f.Pos.Y+float32(f.Offset),
		float32(f.Width),
		float32(f.Height)-float32(f.Offset)+float32(f.cam_fix),
	)
}

func (f *Floor) Render() {
	rl.DrawRectangleRec(
		f.Get_Rec(),
		rl.GetColor(uint(f.Color)),
	)
}
