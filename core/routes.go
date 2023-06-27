package core

import (
	"net/http"
	"fmt"
	)

func (app *Application) routes() *http.ServeMux {
	
	fmt.Println("routes init")
	
	mux := http.NewServeMux()
	
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/users", app.users)
	mux.HandleFunc("/groups", app.groups)
	mux.HandleFunc("/tasks", app.tasks)
	mux.HandleFunc("/stopapp", app.stop)

	mux.HandleFunc("/setappid", app.setAppId)
	mux.HandleFunc("/setappurl", app.setAppToken)
	mux.HandleFunc("/updatestatus", app.updateStatus)
	mux.HandleFunc("/workerstate", app.getWorkerState)
	mux.HandleFunc("/startworker", app.startWorker)
	mux.HandleFunc("/stopworker", app.stopWorker)

	mux.HandleFunc("/updatemyfriends", app.updateMyFriends)
	mux.HandleFunc("/updatemygroups", app.updateMyGroups)
	mux.HandleFunc("/updatemybookmarks", app.updateMyBookmarks)
	mux.HandleFunc("/updategrmembers", app.updateGrMembers)
	mux.HandleFunc("/updateusrfriends", app.updateUsrFriends)
	mux.HandleFunc("/updateusrgroups", app.updateUsrGroups)

	

	return mux
}
