package core

import (
	"fmt"
	"html/template"
	"net/http"
	"vkexplorer/views"
)

func (app *Application) members(w http.ResponseWriter, r *http.Request) {
	groupID := Atoi(r.URL.Path[len("/friends/"):])
	fmt.Println("groupID ID: ", groupID)

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/members.tmpl",
		"./ui/fragments/groupmenu.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/tags.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/userlist.tmpl",
	}

	// get group info
	group := getGroupName(app.db, groupID)

	// get Members list
	page := 0
	pageSize := 10
	members, err := getMembers(app.db, groupID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("members count ", len(members))

	// parsing templates into an internal representation
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// fill
	var t views.NameList
	t.MainMenu = 2
	t.SubMenu = 1
	t.Id = groupID
	t.Name = group
	t.Items = make([]views.NameRec, len(members))
	for i, elem := range members {
		t.Items[i].Id = elem.Uid
		t.Items[i].Name = elem.Name
		fmt.Println(elem.Name)
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
