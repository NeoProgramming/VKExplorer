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

	// arg for pagination
	if args.search != "" {
		t.PageExtraArg = "&search=" + args.search
	}
	if args.sort != "" {
		t.PageExtraArg += "&sort=" + args.sort
		t.PageExtraArg += "&desc=" + Btoa(args.desc)
	}
	if args.filters != "" {
		t.PageExtraArg += "&filters=" + args.filters
	}
	if args.tags != "" {
		t.PageExtraArg += "&tags=" + args.tags
	}

	// arg for sort
	sortExtraArg := fmt.Sprintf("page=%d", args.page)
	if args.search != "" {
		sortExtraArg += "&search=" + args.search
	}
	if args.filters != "" {
		sortExtraArg += "&filters=" + args.filters
	}
	if args.tags != "" {
		sortExtraArg += "&tags=" + args.tags
	}

	// columns
	t.Columns = make([]views.Column, len(args.colunms))
	for i, _ := range t.Columns {
		t.Columns[i].Name = args.colunms[i]
		t.Columns[i].SortExtraArg = &sortExtraArg
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
	args.colunms = []string{"gid", "name", "attrs", "oldest", "newest"}
	t.Items = make([]views.NameRec, len(*groups))
	for i, elem := range *groups {
		t.Items[i].Id = elem.Gid
		t.Items[i].Name = elem.Name
		t.Items[i].Attrs = AttrsToStr(elem.Attrs)
		t.Items[i].Oldest = Tmtoa(elem.Oldest)
		t.Items[i].Newest = Tmtoa(elem.Newest)
	}

	// common & execute
	makeList(w, r, files, args, &t)
}

func (app *Application) makeUserList(w http.ResponseWriter, r *http.Request, files *[]string, users *[]User, args *Args) {
	var t views.NameList

	// user-specific
	args.colunms = []string{"uid", "name", "attrs", "oldest", "newest"}
	t.Items = make([]views.NameRec, len(*users))
	for i, elem := range *users {
		t.Items[i].Id = elem.Uid
		t.Items[i].Name = elem.Name
		t.Items[i].Attrs = AttrsToStr(elem.Attrs)
		t.Items[i].Oldest = Tmtoa(elem.Oldest)
		t.Items[i].Newest = Tmtoa(elem.Newest)
	}

	// common & execute
	makeList(w, r, files, args, &t)
}
