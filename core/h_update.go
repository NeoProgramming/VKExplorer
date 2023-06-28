package core

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// HELPERS

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

func extractId(w http.ResponseWriter, r *http.Request, n string) int {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0
	}
	tid := r.FormValue(n)
	if id, err := strconv.Atoi(tid); err == nil {
		return id
	}
	return 0
}

func extractName(w http.ResponseWriter, r *http.Request, n string) string {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return ""
	}
	tid := r.FormValue(n)
	if strings.HasPrefix(tid, "https://vk.com/") {
		return strings.TrimPrefix(tid, "https://vk.com/")
	}
	return tid
}

// CREATING UPDATE TASKS

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

func (app *Application) updateCheckedGroupMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedGroupMembers")
	gids := extractCheckboxes(w, r)
	for _, gid := range gids {
		app.QueueGroupMembers(gid)
	}
}

func (app *Application) updateCheckedGroupWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedGroupWall")
	gids := extractCheckboxes(w, r)
	for _, gid := range gids {
		app.QueueGroupWall(gid)
	}
}

func (app *Application) updateGroupMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGroupMembers")
	gid := extractName(w, r, "group")
	if gid != "" {
		app.QueueGroupData(gid)
		app.QueueGroupMembers(gid)
	}
}

func (app *Application) updateGroupWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGroupWall")
	gid := extractName(w, r, "group")
	if gid != "" {
		app.QueueGroupData(gid)
		app.QueueGroupWall(gid)
	}
}

func (app *Application) updateGroupInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGroupInfo")
	gid := extractName(w, r, "group")
	if gid != "" {
		app.QueueGroupData(gid)
		app.QueueGroupMembers(gid)
		app.QueueGroupWall(gid)
	}
}

func (app *Application) updateCheckedUserFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedUserFriends")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueUserFriends(uid)
	}
}

func (app *Application) updateCheckedUserGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedUserGroups")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueUserGroups(uid)
	}
}

func (app *Application) updateCheckedUserWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedUserWall")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueUserWall(uid)
	}
}

func (app *Application) updateUserFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserFriends")
	uid := extractName(w, r, "user")
	if uid != "" {
		app.QueueUserData(uid)
		app.QueueUserFriends(uid)
	}
}

func (app *Application) updateUserGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserGroups")
	uid := extractName(w, r, "user")
	if uid != "" {
		app.QueueUserData(uid)
		app.QueueUserGroups(uid)
	}
}

func (app *Application) updateUserWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserWall")
	uid := extractName(w, r, "user")
	if uid != "" {
		app.QueueUserData(uid)
		app.QueueUserWall(uid)
	}
}

func (app *Application) updateUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserInfo")
	uid := extractName(w, r, "user")
	if uid != "" {
		app.QueueUserData(uid)
		app.QueueUserFriends(uid)
		app.QueueUserGroups(uid)
		app.QueueUserWall(uid)
	}
}
