package core

import (
	"html/template"
	"net/http"
	"fmt"
)

// group data
type ViewGroupData struct {
	Gid  int
	Name string
	Title string
}

// Handler for displaying a group content
func (app *Application) group(w http.ResponseWriter, r *http.Request) {
	
	groupID := Atoi(r.URL.Path[len("/group/"):])
    fmt.Printf("Group ID: %d", groupID)

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/group.tmpl",
	}

	// get group info
	group, err := getGroupData(app.db, groupID)
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

	// fill in the list of groups
	var t ViewGroupData	
	t.Name = group

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
