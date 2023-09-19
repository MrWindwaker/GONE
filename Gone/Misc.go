package gone

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	objects "wasm/game/Gone/Objects"
)

func Get_Current_Dir(path string) string {
	pwd, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	p := filepath.Dir(pwd)

	return fmt.Sprintf("%s/%s", p, path)
}

func isLaunchedByDebugger() bool {
	// gops executable must be in the path. See https://github.com/google/gops
	gopsOut, err := exec.Command("gops", strconv.Itoa(os.Getppid())).Output()
	if err == nil && strings.Contains(string(gopsOut), "\\dlv.exe") {
		// our parent process is (probably) the Delve debugger
		return true
	}
	return false
}

var Is_Dev = isLaunchedByDebugger()

func Check_Wall() {
	w := objects.Get_A_Wall()

	fmt.Println("Wall Found:", w)
}
