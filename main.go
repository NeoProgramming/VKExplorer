package main

import (
	"github.com/getlantern/systray"
	"vkexplorer/core"
)

func main() {
	core.InitCore()
	systray.Run(onReady, onExit)
	core.QuitCore()
}

func onReady() {
	core.StartCore()
}

func onExit() {
	// clean up here
}
