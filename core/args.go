package core

import (
	"fmt"
	"net/http"
)

type Args struct {
	id      int    // object id
	name    string // object name
	title   string // page title
	count   int    // total items count (without pagination, filters etc)
	menu    int    // active menu
	submenu int    // active submenu
	page    int    // page number
	size    int    // page size
	search  string
	tags    string
	andor   int
	sort    string
	desc    bool
}

func (args *Args) extractArgs(r *http.Request) {
	// get arguments
	args.size = 10
	args.page = Atodi(r.URL.Query().Get("page"), 1)
	args.search = r.URL.Query().Get("search")
	args.andor = Atoi(r.URL.Query().Get("andor"))
	args.tags = r.URL.Query().Get("tags")
	args.sort = r.URL.Query().Get("sort")
	args.desc = Atoi(r.URL.Query().Get("desc")) != 0

	fmt.Println("search = ", args.search)
	fmt.Println("tags = ", args.tags)
}
