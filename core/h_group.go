package core

import (
	"fmt"
	"html/template"
	"net/http"
	"vkexplorer/views"
)

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

	page := 0
	pageSize := 10

	// get Members list
	members, err := getMembers(app.db, groupID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fill in the list of groups
	var t views.GroupData
	t.Gid = groupID
	t.Name = group
	t.Members = make([]views.UserRec, len(members))
	for i, elem := range members {
		t.Members[i].Uid = elem.Uid
		t.Members[i].Name = elem.Name
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
