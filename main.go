package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/santibarr/cen3031-project/api"
)

func main() {

	//router created
	router := mux.NewRouter()

	// //Err would be 404 not found, Directory does not exist
	// fs := http.FileServer(http.Dir("dist/purfect-partner"))
	// router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	router.HandleFunc("/cats", api.GetCats).Methods("GET")           //read all cats
	router.HandleFunc("/cats-quiz", api.GetCatQuiz).Methods("GET")   // read all cats for quiz
	router.HandleFunc("/cats/{id}", api.UpdateCat).Methods("PUT")    // update cat info
	router.HandleFunc("/cats/{id}", api.DeleteCat).Methods("DELETE") // delete cat
	router.HandleFunc("/cats", api.CreateCat).Methods("POST")        //create a new cat

	//CORS middleware to bypass single origin policy
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"}})

	handler := c.Handler(router)

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
