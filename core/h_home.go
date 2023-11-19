package core

import (
	"html/template"
	"net/http"
)

type ViewHome struct {
	Title       string
	AppID       string
	AppURL      string
	RecentIP    string
	CurrentIP   string
	DBConnected bool
	DBTables    string
	TasksCount  int
	ProxyAddr   string
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {

	// data to pass to the template; any type, reflection in the handler anyway
	data := ViewHome{
		Title:       "This is Users List page",
		AppID:       app.config.AppID,
		AppURL:      app.config.AccessToken,
		RecentIP:    app.config.RecentIP,
		CurrentIP:   GetGlobalIP(),
		DBConnected: app.dbaseConnected,
		DBTables:    GetTables(),
		TasksCount:  getTasksCount(app.db),
		ProxyAddr:   app.config.ProxyAddr,
	}

	// Checks if the current request URL path exactly matches the "/" pattern.
	// If not, it is called http.NotFound() function to return a 404 error to the client.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// We initialize a slice containing paths to two files.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/home.tmpl",
	}

	// Read template file
	// If an error occurs, we will record a detailed error message
	// and send to the user response: 500 Internal Server Error
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Output(2, "parse ok")

	// Write parsed template to HTTP answer body.
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Output(2, "exec ok")

	//w.Write([]byte("Hello from VKExplorer"))
}
