package core

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"vkexplorer/views"
)

// Handler for displaying a list of users
func (app *Application) users(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/users.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/tags.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/userlist.tmpl",
	}

	// pagination
	page := Atodi(r.URL.Query().Get("page"), 1)
	pageSize := 10
	searchStr := r.URL.Query().Get("search")
	andOr := Atoi(r.URL.Query().Get("andor"))
	tagsStr := r.URL.Query().Get("tags")

	fmt.Println("searchStr = ", searchStr)
	fmt.Println("tagsStr = ", tagsStr)

	// get Users list
	users, err := getUsers(app.db, page, pageSize, searchStr, "", false)
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
	var t views.NameList
	t.MainMenu = 1
	t.SubMenu = 0
	t.Items = make([]views.NameRec, len(users))
	for i, elem := range users {
		t.Items[i].Id = elem.Uid
		t.Items[i].Name = elem.Name
		t.Items[i].Oldest = Tmtoa(elem.Oldest)
		t.Items[i].Newest = Tmtoa(elem.Newest)
	}
	t.Title = "Users"
	t.Count = getUsersCount(app.db)
	t.CurrentPage = page
	t.NextPage = page + 1
	t.PrevPage = page - 1
	t.TotalPages = int(math.Ceil(float64(t.Count) / float64(pageSize)))
	t.SearchStr = searchStr
	t.TagsStr = tagsStr
	t.AndOr = andOr
	if searchStr != "" {
		t.SearchArg = "&search=" + searchStr
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
