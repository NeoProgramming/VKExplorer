package core

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/skratchdot/open-golang/open"
	"os/exec"
	"syscall"
)

type TrayItems struct {
	mShow, mHide, mHome, mAbout, mQuit *systray.MenuItem
	mFox                               *systray.MenuItem
}

var Tray TrayItems

func InitTray() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("VKExplorer App")
	systray.SetTooltip("VKExplorer App")

	Tray.mShow = systray.AddMenuItem("Show Console", "Show the console window")
	Tray.mHide = systray.AddMenuItem("Hide Console", "Hide the console window")
	Tray.mHome = systray.AddMenuItem("Open in Default Browser", "Open home page in browser")
	Tray.mFox = systray.AddMenuItem("Open in Firefox", "Open home page in Firefox")
	Tray.mAbout = systray.AddMenuItem("About", "About the app")
	Tray.mQuit = systray.AddMenuItem("Quit", "Quit the whole app")
}

// handle tray menu
func HandleTray() {
	for {
		select {
		case <-Tray.mShow.ClickedCh: // Show console window
			showConsole(true)
		case <-Tray.mHide.ClickedCh: // Hide console window
			showConsole(false)
		case <-Tray.mAbout.ClickedCh: // Show about window
			fmt.Println("VKExplorer program")
		case <-Tray.mHome.ClickedCh: // Open web page
			open.Run("http://127.0.0.1:8080")
		case <-Tray.mFox.ClickedCh: // Open web page in Firefox
			e := exec.Command("C:\\Program Files\\Mozilla Firefox\\firefox.exe", "-new-tab", "http://127.0.0.1:8080").Run()
			fmt.Println(e)
		case <-Tray.mQuit.ClickedCh: // Quit
			systray.Quit()
			return
		}
	}
}
func MyQuit() {
	systray.Quit()
}

func showConsole(show bool) int {
	getWin := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	hwnd, _, _ := getWin.Call()
	if hwnd != 0 {
		showWindowAsync := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindowAsync")
		if show {
			showWindowAsync.Call(hwnd, 1)
		} else {
			showWindowAsync.Call(hwnd, 0)
		}
	}
	return 0
}
