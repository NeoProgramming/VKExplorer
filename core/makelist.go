package core

import (
	"fmt"
	"math"
	"net/http"
	"text/template"
	"vkexplorer/views"
)

func makeList(w http.ResponseWriter, r *http.Request, files *[]string, args *Args, t *views.NameList) {
	// parsing templates into an internal representation
	ts, err := template.ParseFiles(*files...)
	if err != nil {
		App.serverError(w, err)
		return
	}

	// fill Users list
	t.MainMenu = args.menu
	t.SubMenu = args.submenu
	t.Id = args.id
	t.Name = args.name
	t.Title = args.title
	t.Count = args.count
	t.CurrentPage = args.page
	t.NextPage = args.page + 1
	t.PrevPage = args.page - 1
	t.TotalPages = int(math.Ceil(float64(t.Count) / float64(args.size)))
	t.SearchStr = args.search
	t.TagsStr = args.tags
	t.AndOr = args.andor

	//
	if args.search != "" {
		t.SearchArg = "&search=" + args.search
	}
	if args.sort != "" {
		t.SearchArg += "&sort=" + args.sort
		t.SearchArg += "&desc=" + Btoa(args.desc)
	}

	query := fmt.Sprintf("page=%d", args.page)
	if args.search != "" {
		query += "&search="
		query += args.search
	}
	if args.tags != "" {
		query += "&tags="
		query += args.tags
	}

	// columns
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
		App.serverError(w, err)
	}
}

func (app *Application) makeGroupList(w http.ResponseWriter, r *http.Request, files *[]string, groups *[]Group, args *Args) {
	var t views.NameList

	// group-specific
	t.Items = make([]views.NameRec, len(*groups))
	for i, elem := range *groups {
		t.Items[i].Id = elem.Gid
		t.Items[i].Name = elem.Name
		t.Items[i].Oldest = Tmtoa(elem.Oldest)
		t.Items[i].Newest = Tmtoa(elem.Newest)
	}

	// common & execute
	makeList(w, r, files, args, &t)
}

func (app *Application) makeUserList(w http.ResponseWriter, r *http.Request, files *[]string, users *[]User, args *Args) {
	var t views.NameList

	// user-specific
	t.Items = make([]views.NameRec, len(*users))
	for i, elem := range *users {
		t.Items[i].Id = elem.Uid
		t.Items[i].Name = elem.Name
		t.Items[i].Oldest = Tmtoa(elem.Oldest)
		t.Items[i].Newest = Tmtoa(elem.Newest)
	}

	// common & execute
	makeList(w, r, files, args, &t)
}
