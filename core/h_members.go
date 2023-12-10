package core

import (
	"fmt"
	"net/http"
)

func (app *Application) members(w http.ResponseWriter, r *http.Request) {

	// define files
	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/members.tmpl",
		"./ui/fragments/groupmenu.tmpl",
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
	args.id = Atoi(r.URL.Path[len("/friends/"):])
	args.name = getGroupName(app.db, args.id)
	args.menu = 2
	args.submenu = 1
	args.title = "Members"
	args.count = 0

	fmt.Println("groupID ID: ", args.id, " name: ", args.name)

	// get Members list
	members, err := getMembers(app.db, args.id, args.page, args.size, args.search, args.sort, args.desc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// generate web page
	app.makeUserList(w, r, &files, &members, &args)
}
