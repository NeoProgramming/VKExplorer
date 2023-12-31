package core

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
	"vkexplorer/views"
)

type ViewTask struct {
	views.Menu
	Title      string
	Names      []string
	Count      int
	Page       int
	TotalPages int
	PrevPage   int
	NextPage   int
}

// Handler for displaying a list of tasks
func (app *Application) tasks(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/tasks.tmpl",
	}

	// pagination
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

	// get tasks list
	tasks, err := getTasks(app.db, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// template parsing
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// fill out the tasks list
	var t ViewTask
	t.MainMenu = 4
	t.Names = make([]string, len(tasks))
	for i, elem := range tasks {
		t.Names[i] = elem.Name
	}
	t.Title = "Tasks"
	t.Count = getTasksCount(app.db)
	t.Page = page
	t.NextPage = page + 1
	t.PrevPage = page - 1
	t.TotalPages = int(math.Ceil(float64(t.Count) / float64(pageSize)))

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
