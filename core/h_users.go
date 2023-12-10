package core

import (
	"net/http"
)

// Handler for displaying a list of users
func (app *Application) users(w http.ResponseWriter, r *http.Request) {
	// define files
	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/users.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/tags.tmpl",
		"./ui/fragments/filters.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/userlist.tmpl",
		"./ui/fragments/sort.tmpl",
	}

	// get arguments
	var args Args
	args.extractArgs(r)
	args.menu = 1
	args.submenu = 0
	args.title = "Users"
	args.count = getUsersCount(app.db)

	// get Users list
	users, err := getUsers(app.db, args.page, args.size, args.search, args.sort, args.desc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// generate web page
	app.makeUserList(w, r, &files, &users, &args)
}
