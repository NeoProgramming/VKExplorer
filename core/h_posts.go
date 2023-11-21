package core

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
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
	searchStr := r.URL.Query().Get("search")
	fmt.Println("searchStr = ", searchStr)
	andOr := Atoi(r.URL.Query().Get("andor"))
	tagsStr := r.URL.Query().Get("tags")
	fmt.Println("tagsStr = ", tagsStr)

	// get list
	posts, err := getPosts(app.db, page, pageSize, searchStr)
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
	t.Items = make([]views.PostRec, len(posts))
	for i, elem := range posts {
		t.Items[i].Pid = elem.Pid
		t.Items[i].Fid = elem.Fid
	//	t.Items[i].Name = elem.Name
		t.Items[i].Text = elem.Text
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
		t.SearchArg = "&search=" + searchStr
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
