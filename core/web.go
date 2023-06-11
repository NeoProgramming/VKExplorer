package core

import (
	"fmt"
	"log"
	"net/http"
)

func InitWeb() {
	App.srv = &http.Server{
		Addr:     ":8080",
		ErrorLog: App.errorLog,
		Handler:  App.routes(),
	}
	fmt.Println("Web server initialized")
}

func HandleWeb() {
	err := App.srv.ListenAndServe()
	log.Fatal(err)
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, World!"))
		})
		http.ListenAndServe(":8080", nil)
	*/
}
