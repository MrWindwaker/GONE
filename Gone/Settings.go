package gone

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Settings struct {
	Is_Dev bool `json:"isDev"`
}

func Load_Settings() Settings {
	var s Settings

	data_s, err := os.Open("Data/Settings.json")
	if err != nil {
		fmt.Println("Settings.josn not Found")
	}

	defer data_s.Close()

	setting_d, err := io.ReadAll(data_s)
	if err != nil {
		fmt.Println("Error reading Settings data")
	}

	json.Unmarshal(setting_d, &s)

	return s
}
