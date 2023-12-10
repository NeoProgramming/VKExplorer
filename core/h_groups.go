package core

import (
	"net/http"
)

// Handler for displaying a list of groups
func (app *Application) groups(w http.ResponseWriter, r *http.Request) {
	// define files
	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/groups.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/tags.tmpl",
		"./ui/fragments/filters.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/grouplist.tmpl",
		"./ui/fragments/sort.tmpl",
	}

	// get arguments
	var args Args
	args.extractArgs(r)
	args.menu = 2
	args.submenu = 0
	args.id = 0
	args.title = "Groups"
	args.count = getGroupsCount(app.db)

	// get list
	groups, err := getGroups(app.db, args.page, args.size, args.search, args.sort, args.desc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// generate web page
	app.makeGroupList(w, r, &files, &groups, &args)
}
