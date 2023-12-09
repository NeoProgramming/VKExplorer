package core

import (
	"fmt"
	"net/http"
)

// Handler for displaying a group content
func (app *Application) friends(w http.ResponseWriter, r *http.Request) {
	// define files
	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/friends.tmpl",
		"./ui/fragments/usermenu.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/tags.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/userlist.tmpl",
		"./ui/fragments/sort.tmpl",
	}

	// get arguments
	var args Args
	args.extractArgs(r)
	args.id = Atoi(r.URL.Path[len("/friends/"):])
	args.name = getUserName(app.db, args.id)
	args.menu = 1
	args.submenu = 1
	args.title = "Friends"
	args.count = 0

	fmt.Println("User ID: ", args.id, " name: ", args.name)

	// get Friends list
	friends, err := getFriends(app.db, args.id, args.page, args.size, args.search, args.sort, args.desc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// generate web page
	app.makeUserList(w, r, &files, &friends, &args)
}
