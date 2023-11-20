package core

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"vkexplorer/views"
)

func (app *Application) groupwall(w http.ResponseWriter, r *http.Request) {
	groupID := Atoi(r.URL.Path[len("/groupwall/"):])
	fmt.Println("groupID ID: ", groupID)

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/groupwall.tmpl",
		"./ui/fragments/groupmenu.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/postlist.tmpl",
	}

	// get group info
	group, err := getGroupData(app.db, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get Wall
	page := 0
	pageSize := 10
	wall, err := getWall(app.db, -groupID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("wall records count = ", len(wall))

	// parsing templates into an internal representation
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// fill
	var t views.PostsList
	t.Id = groupID
	t.Name = group
	t.Items = make([]views.PostRec, len(wall))
	for i, elem := range wall {
		t.Items[i].Pid = elem.Pid
		t.Items[i].Fid = elem.Fid
		if elem.Name == "" {
			t.Items[i].Name = "! " + strconv.Itoa(elem.Fid)
		} else {
			t.Items[i].Name = elem.Name
		}
		t.Items[i].Text = elem.Text
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
