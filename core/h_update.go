package core

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	return 0
}

// 

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
	fmt.Println("updateGroupMembers")
	gids := extractCheckboxes(w, r)
	for _, gid := range gids {
		app.QueueGroupMembers(gid)
	}
}

func (app *Application) updateCheckedGroupWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGroupWall")
	gids := extractCheckboxes(w, r)
	for _, gid := range gids {
		app.QueueGroupWall(gid)
	}
}

func (app *Application) updateGroupMembers(w http.ResponseWriter, r *http.Request) {
	
}

func (app *Application) updateGroupWall(w http.ResponseWriter, r *http.Request) {
	
}

func (app *Application) updateGroupInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGroupInfo")
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	groupId := r.FormValue("group")
	fmt.Println(groupId)
	if id, err := strconv.Atoi(groupId); err == nil {
		app.QueueGroupMembers(id)
		app.QueueGroupWall(id)
	}
}

func (app *Application) updateCheckedUserFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserFriends")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueUserFriends(uid)
	}
}

func (app *Application) updateCheckedUserGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserGroups")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueUserGroups(uid)
	}
}

func (app *Application) updateCheckedUserWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUsrWall")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueUserWall(uid)
	}
}

func (app *Application) updateUserFriends(w http.ResponseWriter, r *http.Request) {
	
}

func (app *Application) updateUserGroups(w http.ResponseWriter, r *http.Request) {
	
}

func (app *Application) updateUserWall(w http.ResponseWriter, r *http.Request) {
	
}

func (app *Application) updateUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserInfo")
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId := r.FormValue("user")
	fmt.Println(userId)
	if id, err := strconv.Atoi(userId); err == nil {
		app.QueueUserFriends(id)
		app.QueueUserGroups(id)
		app.QueueUserWall(id)
	}
}
