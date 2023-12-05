package core

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"text/template"
	"vkexplorer/views"
)

// Handler for displaying a list of groups
func (app *Application) groups(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/groups.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/tags.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/grouplist.tmpl",
		"./ui/fragments/sort.tmpl",
	}

	// pagination: we take the page number from the URL, 1 by default
	page := Atodi(r.URL.Query().Get("page"), 1)
	pageSize := 10
	search := r.URL.Query().Get("search")
	andor := Atoi(r.URL.Query().Get("andor"))
	tags := r.URL.Query().Get("tags")
	sort := r.URL.Query().Get("sort")
	desc := Atoi(r.URL.Query().Get("desc"))

	fmt.Println("search = ", search)
	fmt.Println("tags = ", tags)

	// get list
	groups, err := getGroups(app.db, page, pageSize, search, sort, desc != 0)
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
	var t views.NameList
	t.MainMenu = 2
	t.SubMenu = 0

	t.Items = make([]views.NameRec, len(groups))
	for i, elem := range groups {
		t.Items[i].Id = elem.Gid
		t.Items[i].Name = elem.Name
		t.Items[i].Oldest = Tmtoa(elem.Oldest)
		t.Items[i].Newest = Tmtoa(elem.Newest)
	}
	t.Title = "Groups"
	t.Count = getGroupsCount(app.db)
	t.CurrentPage = page
	t.NextPage = page + 1
	t.PrevPage = page - 1
	t.TotalPages = int(math.Ceil(float64(t.Count) / float64(pageSize)))
	t.SearchStr = search
	t.TagsStr = tags
	t.AndOr = andor
	if search != "" {
		t.SearchArg += "&search=" + search
	}
	if sort != "" {
		t.SearchArg += "&sort=" + sort
		t.SearchArg += "&desc=" + strconv.Itoa(desc)
	}

	query := fmt.Sprintf("page=%d", page)
	if search != "" {
		query += "&search="
		query += search
	}
	if tags != "" {
		query += "&tags="
		query += tags
	}

	t.Columns = make([]views.Column, 4)
	t.Columns[0].Name = "gid"
	t.Columns[0].Title = "URL"
	t.Columns[1].Name = "name"
	t.Columns[1].Title = "Name"
	t.Columns[2].Title = "Oldest"
	t.Columns[2].Name = "oldest"
	t.Columns[3].Title = "Newest"
	t.Columns[3].Name = "newest"
	for i, _ := range t.Columns {
		t.Columns[i].Query = &query
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
