package gone

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	objs "wasm/game/Gone/Objects"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func load_json_floors() objs.FloorsJson {
	var f objs.FloorsJson

	data_f, err := os.Open("Data/Floor.json")
	if err != nil {
		fmt.Println("Floor.json Not Found")
	}

	defer data_f.Close()

	floor_data, err := io.ReadAll(data_f)
	if err != nil {
		fmt.Println("Error reading Floor Data")
	}

	json.Unmarshal(floor_data, &f)

	return f
}

func create_floor() map[string]objs.Floor {
	fs := load_json_floors()
	floor := make(map[string]objs.Floor)

	for _, f := range fs.Floors {
		floor[f.Name] = objs.New_Floor(
			rl.NewVector2(float32(f.Pos.X), float32(f.Pos.Y)),
			f.Width,
			f.Height,
			f.Offset,
		)
	}

	return floor
}
