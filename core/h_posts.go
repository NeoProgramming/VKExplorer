package core

import (
	"fmt"
	"math"
	"net/http"
	"text/template"
	"time"
	"vkexplorer/views"
)

// Handler for displaying a list of groups
func (app *Application) posts(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/posts.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/tags.tmpl",
		"./ui/fragments/pagination.tmpl",
		"./ui/fragments/postlist.tmpl",
	}

	// pagination: we take the page number from the URL, 1 by default
	page := Atodi(r.URL.Query().Get("page"), 1)
	pageSize := 10
	searchStr := r.URL.Query().Get("search")
	andOr := Atoi(r.URL.Query().Get("andor"))
	tagsStr := r.URL.Query().Get("tags")

	fmt.Println("page = ", page)
	fmt.Println("search = ", searchStr)
	fmt.Println("tags = ", tagsStr)

	// get list
	posts, err := getPosts(app.db, page, pageSize, searchStr, "", false)
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

	// fill in the list of posts
	var t views.PostsList
	t.MainMenu = 3
	t.Items = make([]views.PostRec, len(posts))
	for i, elem := range posts {
		t.Items[i].Pid = elem.Pid
		t.Items[i].Fid = elem.Fid
		//	t.Items[i].Name = elem.Name
		t.Items[i].Text = elem.Text
		t.Items[i].Date = Ttoa(time.Unix(elem.Date, 0))
		//	t.Items[i].UpdateTime = elem.UpdatedAt.Format("2006-01-02 15:04:05")
	}
	t.Title = "Posts"
	t.Count = getPostsCount(app.db)
	t.CurrentPage = page
	t.NextPage = page + 1
	t.PrevPage = page - 1
	t.TotalPages = int(math.Ceil(float64(t.Count) / float64(pageSize)))
	t.SearchStr = searchStr
	t.TagsStr = tagsStr
	t.AndOr = andOr
	if searchStr != "" {
		t.PageExtraArg = "&search=" + searchStr
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
