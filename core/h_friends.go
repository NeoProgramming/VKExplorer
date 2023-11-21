package core

import (
	"fmt"
	"html/template"
	"net/http"
	"vkexplorer/views"
)

// Handler for displaying a group content
func (app *Application) friends(w http.ResponseWriter, r *http.Request) {

	userID := Atoi(r.URL.Path[len("/friends/"):])
	fmt.Println("User ID: ", userID)

	files := []string{
		"./ui/pages/base.tmpl",
		"./ui/pages/friends.tmpl",
		"./ui/fragments/usermenu.tmpl",
		"./ui/fragments/search.tmpl",
		"./ui/fragments/userlist.tmpl",
	}

	// get User data
	user, err := getUserData(app.db, userID)
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

	page := 0
	pageSize := 10

	// get Friends list
	friends, err := getFriends(app.db, userID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("frinds count ", len(friends))
	
	// fill 
	var t views.NameList
	t.Id = userID
	t.Name = user
	t.Items = make([]views.NameRec, len(friends))
	for i, elem := range friends {
		t.Items[i].Id = elem.Uid
		t.Items[i].Name = elem.Name
		fmt.Println(elem.Name)
	}

	// execute templates
	err = ts.Execute(w, t)
	if err != nil {
		app.serverError(w, err)
	}
}
