package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maotterson/sample-go-api/controllers"
)

func main()  {
	fmt.Println("Rest API Running")
	setupRoutes()
}

func setupRoutes()  {
	router := mux.NewRouter()

	router.HandleFunc("/", routeHome)

	// articles
	router.HandleFunc("/api/articles",controllers.GetArticles).Methods("GET")
	router.HandleFunc("/api/articles/{id}",controllers.GetArticle).Methods("GET")
	router.HandleFunc("/api/articles",controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/api/articles/{id}",controllers.EditArticle).Methods("PUT")
	router.HandleFunc("/api/articles/{id}",controllers.DeleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func routeHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home page")
	fmt.Println("Endpoint Hit: Request received @ home")
}
