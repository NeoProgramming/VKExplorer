package core

import (
	"syscall"
)

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
