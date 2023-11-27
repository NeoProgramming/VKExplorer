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
		"./ui/pages/base.tmpl",
		"./ui/pages/group.tmpl",
		"./ui/fragments/groupmenu.tmpl",
	}

	// get group info
	group, err := getGroupInfo(app.db, groupID)
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

	// fill GroupData
	var t views.GroupData
	t.MainMenu = 2
	t.SubMenu = 0
	t.Id = groupID
	t.Name = group.Name
	t.MembersUpdated = group.MembersUpdated.Format("06-01-02 15:04")
	t.WallUpdated = group.WallUpdated.Format("06-01-02 15:04")

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
