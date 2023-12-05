package core

import (
	"fmt"
	"math"
	"net/http"
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
	searchStr := r.URL.Query().Get("search")
	andOr := Atoi(r.URL.Query().Get("andor"))
	tagsStr := r.URL.Query().Get("tags")
	sort := r.URL.Query().Get("sort")
	desc := Atoi(r.URL.Query().Get("desc"))

	fmt.Println("searchStr = ", searchStr)
	fmt.Println("tagsStr = ", tagsStr)

	// get list
	groups, err := getGroups(app.db, page, pageSize, searchStr, sort, desc!=0)
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
		oldest := minTime(elem.MembersUpdated, elem.WallUpdated)
		newest := maxTime(elem.MembersUpdated, elem.WallUpdated)
		t.Items[i].OldestUpdateTime = Tmtoa(oldest)
		t.Items[i].NewestUpdateTime = Tmtoa(newest)
	}
	t.Title = "Groups"
	t.Count = getGroupsCount(app.db)
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
	
	query := fmt.Sprintf("page=%d", page)
	if searchStr != "" {
		query += "&search="
		query += searchStr
	}
	if tagsStr != "" {
		query += "&tags="
		query += tagsStr
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
