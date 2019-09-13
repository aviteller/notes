package main

import (
	"fmt"
	"log"
	"net/http"

	"./controllers"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},                                               // All origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"}, // Allowing only get, just an example
	})

	router.HandleFunc("/api/notes", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/api/notes", controllers.GetNotes).Methods("GET")
	// router.HandleFunc("/api/birthday/{id}", controllers.GetBirthday).Methods("GET")
	router.HandleFunc("/api/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	router.HandleFunc("/api/notes/{id}", controllers.UpdateNote).Methods("PUT")

	// router.Use(app.JwtAuthentication)
	// tmpl := template.Must(template.ParseFiles("static/public/index.html"))
	// fs := http.FileServer(http.Dir("./static"))
	// router.PathPrefix("/public/").Handler(fs)
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl.Execute(w, 0)
	// })

	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static/public").HTTPBox()))

	port := "9000"
	if port == "" {
		port = "8000"
	}

	err2 := open.Start("http://localhost:9000")
	if err2 != nil {
		log.Println(err2)
	}

	err := http.ListenAndServe(":"+port, c.Handler(router))
	if err != nil {
		fmt.Println(err)
	}

}
