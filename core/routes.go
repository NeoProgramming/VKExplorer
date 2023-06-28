package core

import (
	"fmt"
	"net/http"
)

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/images/favicon.ico")
}

func (app *Application) routes() *http.ServeMux {

	fmt.Println("routes init")

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/favicon.ico", faviconHandler)

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/users", app.users)
	mux.HandleFunc("/groups", app.groups)
	mux.HandleFunc("/tasks", app.tasks)
	mux.HandleFunc("/exit", app.exit)

	mux.HandleFunc("/set-app-id", app.setAppId)
	mux.HandleFunc("/set-app-url", app.setAppToken)
	mux.HandleFunc("/get-server-status", app.getServerStatus)
	mux.HandleFunc("/get-worker-status", app.getWorkerStatus)
	mux.HandleFunc("/start-worker", app.startWorker)
	mux.HandleFunc("/stop-worker", app.stopWorker)

	mux.HandleFunc("/update-my-friends", app.updateMyFriends)
	mux.HandleFunc("/update-my-groups", app.updateMyGroups)
	mux.HandleFunc("/update-my-bookmarks", app.updateMyBookmarks)
	
	mux.HandleFunc("/update-group-members", app.updateGroupMembers)
	mux.HandleFunc("/update-group-wall", app.updateGroupWall)
	mux.HandleFunc("/update-group-info", app.updateGroupInfo)
	
	mux.HandleFunc("/update-checked-group-members", app.updateCheckedGroupMembers)
	mux.HandleFunc("/update-checked-group-wall", app.updateCheckedGroupWall)
	
	mux.HandleFunc("/update-user-friends", app.updateUserFriends)
	mux.HandleFunc("/update-user-groups", app.updateUserGroups)
	mux.HandleFunc("/update-user-wall", app.updateUserWall)
	mux.HandleFunc("/update-user-info", app.updateUserInfo)	
	
	mux.HandleFunc("/update-checked-user-friends", app.updateCheckedUserFriends)
	mux.HandleFunc("/update-checked-user-groups", app.updateCheckedUserGroups)
	mux.HandleFunc("/update-checked-user-wall", app.updateCheckedUserWall)

	return mux
}
