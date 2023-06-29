package core

import (
	"html/template"
	"net/http"
	"fmt"
)

type ViewUserData struct {
	Uid  int
	Name string
	Title string
}

// Handler for displaying a user page
func (app *Application) user(w http.ResponseWriter, r *http.Request) {
	
	userID := Atoi(r.URL.Path[len("/user/"):])
    fmt.Printf("User ID: %d", userID)

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

	// fill Users list
	var t ViewUserData
	t.Name = user
	
	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
