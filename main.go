package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santibarr/cen3031-project/api"
)

func main() {

	//router created

	//slice of cats
	// var cats []scrape.Cat

	// cats = append(cats, scrape.MiamiDade(cats)...)
	// cats = append(cats, scrape.LakeCounty(cats)...)
	// //cats = append(cats, scrape.Peggy(cats)...)
	// //cats = append(cats, scrape.Marathon(cats)...)
	// //cats = append(cats, scrape.KeyWest(cats)...)
	router := mux.NewRouter()
	//println(len(cats))
	//edit

	//Err would be 404 not found, Directory does not exist
	fs := http.FileServer(http.Dir("dist/purfect-partner"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	router.HandleFunc("/cats", api.GetCats).Methods("GET")           //read all cats
	router.HandleFunc("/cats-quiz", api.GetCatQuiz).Methods("GET")       // read cat quiz
	router.HandleFunc("/cats/{id}", api.UpdateCat).Methods("PUT")    // update cat info
	router.HandleFunc("/cats/{id}", api.DeleteCat).Methods("DELETE") // delete cat
	router.HandleFunc("/cats", api.CreateUser).Methods("POST")        //create a new cat

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
