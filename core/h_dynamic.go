package core

import (
	"fmt"
	"net/http"
	"time"
)

// AJAX - setting the application ID
func (app *Application) setAppId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("setAppId")
	if r.Method == http.MethodPost {
		app.config.AppID = r.FormValue("app_id")
		fmt.Println("app_id:", app.config.AppID)
		// ... update user information in the database ...
		w.Write([]byte("setAppId ok"))
		SaveConfig()
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// AJAX - setting the access token
func (app *Application) setAppToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("setAppToken")
	if r.Method == http.MethodPost {
		urlStr := r.FormValue("app_url")
		if urlStr != "" {
			app.config.AccessToken = extractAccessToken(urlStr)
			app.config.RecentIP = GetGlobalIP()
		}
		fmt.Println("app_token:", app.config.AccessToken)
		// ... update user information in the database ...
		w.Write([]byte("setAppToken ok"))
		SaveConfig()
		InitVK()
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

//  SSE - get server status
func (app *Application) getServerStatus(w http.ResponseWriter, r *http.Request) {
	// Set response headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for {
		// Get status of goroutine and send to client
		status := app.GetStatus()
		fmt.Fprintf(w, "data: %s\n\n", status)
		w.(http.Flusher).Flush()

		time.Sleep(time.Second)

		// Check for SSE connection closure
		select {
		case <-r.Context().Done():
			fmt.Println("SSE connection closed")
			return
		default:
		}
	}
}

// AJAX - getting the state of the worker goroutine
func (app *Application) getWorkerStatus(w http.ResponseWriter, r *http.Request) {
	if app.running {
		fmt.Fprint(w, "true")
	} else {
		fmt.Fprint(w, "false")
	}
}

// AJAX - running a worker goroutine
func (app *Application) startWorker(w http.ResponseWriter, r *http.Request) {
	if app.vk == nil {
		InitVK()
	}
	if !app.running {
		app.running = true
		app.wg.Add(1)
		go app.worker()
		fmt.Fprint(w, "true")
		fmt.Println("Worker STARTED")
	}
}

// AJAX - stopping a worker goroutine
func (app *Application) stopWorker(w http.ResponseWriter, r *http.Request) {
	if app.running {
		app.running = false
		app.wg.Wait()
		fmt.Fprint(w, "false")
		fmt.Println("Worker stopped manually")
	}
}
