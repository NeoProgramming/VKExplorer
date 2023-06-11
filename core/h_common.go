package core

import (
	"fmt"
	"github.com/getlantern/systray"
	"net/http"
)

// handlers are registered in routes.go

// stop handler
func (app *Application) stop(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(200)
		fmt.Fprint(w, "Shutting down the server...")
		//fmt.Fprint(w, "Reload main page: http://127.0.0.1:8080")
		go func() {
			systray.Quit()
			fmt.Fprintln(w, "Server stopped.")
		}()
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
