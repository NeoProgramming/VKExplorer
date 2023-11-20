package core

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
	"vkexplorer/views"
)

// Handler for displaying a list of users
func (app *Application) users(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/users.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/userlist.tmpl",
	}

	// pagination
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

	// get Users list
	users, err := getUsers(app.db, page, pageSize)
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
	var t views.UsersList
	t.Items = make([]views.UserRec, len(users))
	for i, elem := range users {
		t.Items[i].Id = elem.Uid
		t.Items[i].Name = elem.Name
		t.Items[i].UpdateTime = elem.UpdatedAt.Format("2006-01-02 15:04:05")
	}
	t.Title = "Users"
	t.Count = getUsersCount(app.db)
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
