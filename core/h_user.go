package core

import (
	"fmt"
	"html/template"
	"net/http"
	"vkexplorer/views"
)

// Handler for displaying a user page
func (app *Application) user(w http.ResponseWriter, r *http.Request) {

	userID := Atoi(r.URL.Path[len("/user/"):])
	fmt.Printf("User ID: %d\n", userID)

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/user.tmpl",
	}

	// get User data
	user, err := getUserData(app.db, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// parsing templates into an internal representation
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	page := 0
	pageSize := 10

	// get Friends list
	friends, err := getFriends(app.db, userID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get Groups list
	groups, err := getMemberships(app.db, userID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fill UserData
	var t views.UserData
	t.Uid = userID
	t.Name = user
	t.Friends = make([]views.UserRec, len(friends))
	for i, elem := range friends {
		t.Friends[i].Uid = elem.Uid
		t.Friends[i].Name = elem.Name
	}
	t.Groups = make([]views.GroupRec, len(groups))
	for i, elem := range groups {
		t.Groups[i].Gid = elem.Gid
		t.Groups[i].Name = elem.Name
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
