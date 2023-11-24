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
		"./ui/pages/base.tmpl",
		"./ui/pages/user.tmpl",
		"./ui/fragments/usermenu.tmpl",
	}

	// get user info
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

	// fill UserData
	var t views.UserData
	t.MainMenu = 1
	t.SubMenu = 0	
	t.Id = userID
	t.Name = user

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
