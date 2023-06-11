package core

import (
	"html/template"
	"net/http"
	"strconv"
)

type ViewTask struct {
	Title string
	Names []string
}

// Handler for displaying a list of tasks
func (app *Application) tasks(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/tasks.tmpl",
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
	t.Names = make([]string, len(tasks))
	for i, elem := range tasks {
		t.Names[i] = elem.Name
	}
	t.Title = "Tasks"

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
