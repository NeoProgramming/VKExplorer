package core

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

//  SSE - status update
func (app *Application) updateStatus(w http.ResponseWriter, r *http.Request) {
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
func (app *Application) getWorkerState(w http.ResponseWriter, r *http.Request) {
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

// AJAX - setting the task "update my friends list"
func (app *Application) updateMyFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateMyFriends")
	app.QueueMyFriends()
}

// AJAX - setting the task "update list of my groups"
func (app *Application) updateMyGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateMyGroups")
	app.QueueMyGroups()
}

// AJAX - setting the task  "update list of my bookmarks"
func (app *Application) updateMyBookmarks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateMyBookmarks")
	app.QueueMyBookmarks()
}

func extractCheckboxes(w http.ResponseWriter, r *http.Request) []int {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return make([]int, 0)
	}

	// Get the values of the checkboxes
	checkboxValues := r.FormValue("checkbox")
	fmt.Println(checkboxValues)

	// Split the comma-separated values into a slice
	values := strings.Split(checkboxValues, ",")

	ints := make([]int, 0, len(values))
	for _, token := range values {
		if i, err := strconv.Atoi(token); err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func (app *Application) updateGrMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGrMembers")
	gids := extractCheckboxes(w, r)
	for _, gid := range gids {
		app.QueueGroupMembers(gid)
	}
}

func (app *Application) updateUsrFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUsrFriends")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueUserFriends(uid)
	}
}

func (app *Application) updateUsrGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUsrGroups")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueUserGroups(uid)
	}
}

func (app *Application) updateUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserInfo")
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId := r.FormValue("user")
	fmt.Println(userId)
	//if i, err := strconv.Atoi(userId); err == nil {
	app.QueueUser(userId)
	//}
}

func (app *Application) updateGroupInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGroupInfo")
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	groupId := r.FormValue("group")
	fmt.Println(groupId)
	//if i, err := strconv.Atoi(groupId); err == nil {
	app.QueueGroup(groupId)
	//}
}
