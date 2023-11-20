package core

import (
	"fmt"
	"html/template"
	"net/http"
	"vkexplorer/views"
)

// Handler for displaying a user page
func (app *Application) membership(w http.ResponseWriter, r *http.Request) {

	userID := Atoi(r.URL.Path[len("/membership/"):])
	fmt.Printf("User ID: %d\n", userID)

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/membership.tmpl",
		"./ui/fragments/usermenu.tmpl",
		"./ui/fragments/grouplist.tmpl",
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

	// get Groups list
	groups, err := getMemberships(app.db, userID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fill UserData
	var t views.GroupsList
	t.Id = userID
	t.Name = user
	t.Items = make([]views.GroupRec, len(groups))
	for i, elem := range groups {
		t.Items[i].Id = elem.Gid
		t.Items[i].Name = elem.Name
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
