package main

import (
	"blackjack/gameplay"
	"blackjack/ui"
	ebitenui "blackjack/ui/ebitenUI"
)

func main() {
	// ui mode
	gameplay := new(gameplay.SinglePlayer)
	ebitenUI := new(ebitenui.EbitenUI)
	var ui ui.UI = ebitenUI
	go gameplay.Init(&ui)
	ebitenUI.Init()

	// console mode
	//gameplay := new(gameplay.SinglePlayer)
	//consoleUI := new(ui.ConsoleUi)
	//var ui ui.UI = consoleUI
	//gameplay.Init(&ui)
}
