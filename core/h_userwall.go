package core

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"vkexplorer/views"
)

func (app *Application) userwall(w http.ResponseWriter, r *http.Request) {
	userID := Atoi(r.URL.Path[len("/userwall/"):])
	fmt.Println("userID: ", userID)

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/userwall.tmpl",
		"./ui/fragments/usermenu.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/tags.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/postlist.tmpl",
	}

	// get user info
	user, err := getUserData(app.db, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get Wall
	page := 0
	pageSize := 10
	wall, err := getWall(app.db, userID, page, pageSize)
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
	t.Id = userID
	t.Name = user
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
