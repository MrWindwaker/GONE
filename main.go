package main

import eng "wasm/game/Gone"

func main() {
	e := eng.Get_Engine()

	e.Init()

	for !e.Should_Close {
		e.Update()
		e.Render()
	}

	e.Close()
}
