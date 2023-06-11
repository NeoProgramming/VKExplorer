package core

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type ViewUser struct {
	Uid  int
	Name string
}

type ViewUsersList struct {
	Title       string
	Items       []ViewUser
	Count       int
	CurrentPage int
	TotalPages  int
	PrevPage    int
	NextPage    int
}

// Handler for displaying a list of users
func (app *Application) users(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/users.tmpl",
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
	var t ViewUsersList
	t.Items = make([]ViewUser, len(users))
	for i, elem := range users {
		t.Items[i].Uid = elem.Uid
		t.Items[i].Name = elem.Name
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
