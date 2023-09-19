package objects

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Wall struct {
	Pos    rl.Vector2 `json:"pos"`
	Width  int        `json:"width"`
	Height int        `json:"height"`
}

func Get_A_Wall() Wall {
	var w Wall

	data_w, err := os.Open("Data/Walls.json")
	if err != nil {
		fmt.Println("Walls.json not Found")
	}

	defer data_w.Close()

	wall_d, err := io.ReadAll(data_w)
	if err != nil {
		fmt.Println("Error reading Wall data")
	}

	json.Unmarshal(wall_d, &w)

	return w
}
