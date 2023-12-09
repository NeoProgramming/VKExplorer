package core

import (
	"fmt"
	"net/http"
)

// Handler for displaying a user page
func (app *Application) membership(w http.ResponseWriter, r *http.Request) {

	// define files
	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/membership.tmpl",
		"./ui/fragments/usermenu.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/tags.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/grouplist.tmpl",
		"./ui/fragments/sort.tmpl",
	}

	// get arguments
	var args Args
	args.extractArgs(r)
	args.id = Atoi(r.URL.Path[len("/membership/"):])
	args.name = getUserName(app.db, args.id)
	args.menu = 1
	args.submenu = 2
	args.title = "Membership"

	// get Groups list
	groups, err := getMemberships(app.db, args.id, args.page, args.size, args.search, args.sort, args.desc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("User ID: %d\n", args.id)

	// generate web page
	app.makeGroupList(w, r, &files, &groups, &args)
}
