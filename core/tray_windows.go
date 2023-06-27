package core

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/skratchdot/open-golang/open"
	"os/exec"
)

func InitTray() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("VKExplorer App")
	systray.SetTooltip("VKExplorer App")

	App.mShow = systray.AddMenuItem("Show Console", "Show the console window")
	App.mHide = systray.AddMenuItem("Hide Console", "Hide the console window")
	App.mHome = systray.AddMenuItem("Open in Default Browser", "Open home page in browser")
	App.mFox = systray.AddMenuItem("Open in Firefox", "Open home page in Firefox")
	App.mAbout = systray.AddMenuItem("About", "About the app")
	App.mQuit = systray.AddMenuItem("Quit", "Quit the whole app")
}

// handle tray menu
func HandleTray() {
	for {
		select {
		case <-App.mShow.ClickedCh: // Show console window
			showConsole(true)
		case <-App.mHide.ClickedCh: // Hide console window
			showConsole(false)
		case <-App.mAbout.ClickedCh: // Show about window
			fmt.Println("VKExplorer program")
		case <-App.mHome.ClickedCh: // Open web page
			open.Run("http://127.0.0.1:8080")
		case <-App.mFox.ClickedCh: // Open web page in Firefox
			e := exec.Command("C:\\Program Files\\Mozilla Firefox\\firefox.exe", "-new-tab", "http://127.0.0.1:8080").Run()
			fmt.Println(e)
		case <-App.mQuit.ClickedCh: // Quit
			systray.Quit()
			return
		}
	}
}
