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
		fmt.Println("extractName: http error")
		return ""
	}
	tid := r.FormValue(n)
	if strings.HasPrefix(tid, "https://vk.com/") {
		tid = strings.TrimPrefix(tid, "https://vk.com/")
		fmt.Println("extractName: url, name = " + tid)
		return tid
	}
	fmt.Println("extractName: non-url, name = " + tid)
	return tid
}

// CREATING UPDATE TASKS

// AJAX - setting the task "update my friends list"
func (app *Application) updateMyFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateMyFriends")
	app.QueueByType(TT_MyFriends, "MyFriends")
}

// AJAX - setting the task "update list of my groups"
func (app *Application) updateMyGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateMyGroups")
	app.QueueByType(TT_MyGroups, "MyGroups")
}

// AJAX - setting the task  "update list of my bookmarks"
func (app *Application) updateMyBookmarks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateMyBookmarks")
	app.QueueByType(TT_MyBookmarks, "MyBookmarks")
}

func (app *Application) updateCheckedGroupMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedGroupMembers")
	gids := extractCheckboxes(w, r)
	for _, gid := range gids {
		app.QueueById(TT_GroupMembers, gid, "GroupMembers")
	}
}

func (app *Application) updateCheckedGroupWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedGroupWall")
	gids := extractCheckboxes(w, r)
	for _, gid := range gids {
		app.QueueById(TT_GroupWall, gid, "GroupWall")
	}
}

func (app *Application) updateGroupMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGroupMembers")
	gid := extractName(w, r, "group")
	if gid != "" {
		app.QueueByName(TT_GroupMembersByName, gid)
	}
}

func (app *Application) updateGroupWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGroupWall")
	gid := extractName(w, r, "group")
	if gid != "" {
		app.QueueByName(TT_GroupWallByName, gid)
	}
}

func (app *Application) updateGroupData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateGroupData")
	gid := extractName(w, r, "group")
	if gid != "" {
		app.QueueByName(TT_GroupDataByName, gid)
	}
}

func (app *Application) updateCheckedUserFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedUserFriends")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueById(TT_UserFriends, uid, "UserFriends")
	}
}

func (app *Application) updateCheckedUserGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedUserGroups")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueById(TT_UserGroups, uid, "UserGroups")
	}
}

func (app *Application) updateCheckedUserWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateCheckedUserWall")
	uids := extractCheckboxes(w, r)
	for _, uid := range uids {
		app.QueueById(TT_UserWall, uid, "UserWall")
	}
}

func (app *Application) updateUserFriends(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserFriends")
	uid := extractName(w, r, "user")
	if uid != "" {
		app.QueueByName(TT_UserFriendsByName, uid)
	}
}

func (app *Application) updateUserGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserGroups")
	uid := extractName(w, r, "user")
	if uid != "" {
		app.QueueByName(TT_UserGroupsByName, uid)
	}
}

func (app *Application) updateUserWall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserWall")
	uid := extractName(w, r, "user")
	if uid != "" {
		app.QueueByName(TT_UserWallByName, uid)
	}
}

func (app *Application) updateUserData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateUserData")
	uid := extractName(w, r, "user")
	if uid != "" {
		app.QueueByName(TT_UserDataByName, uid)
	}
}
