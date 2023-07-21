package core

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
	"vkexplorer/views"
)

// Handler for displaying a list of groups
func (app *Application) groups(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/groups.tmpl",
	}

	// pagination: we take the page number from the URL, 1 by default
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	pageSize := 10

	// get list
	groups, err := getGroups(app.db, page, pageSize)
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
	var t views.GroupsList
	t.Items = make([]views.GroupRec, len(groups))
	for i, elem := range groups {
		t.Items[i].Gid = elem.Gid
		t.Items[i].Name = elem.Name
	}
	t.Title = "Groups"
	t.Count = getGroupsCount(app.db)
	t.CurrentPage = page
	t.NextPage = page + 1
	t.PrevPage = page - 1
	t.TotalPages = int(math.Ceil(float64(t.Count) / float64(pageSize)))

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
