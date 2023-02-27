package main

//import packages - updated
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//router created
	router := mux.NewRouter()

	//this should work
	fs := http.FileServer(http.Dir("src"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	//example for now, but should invoke scraper.go to get cat objects.
	cats = append(cats, Cat{Name: "Smithy", Age: "1 year old", Gender: "Male", Loc: "Gainesville Cat Cafe", ID: "A245676"})
	cats = append(cats, Cat{Name: "Ruby", Age: " 4 months old", Gender: "Female", Loc: "Miami Animal Shelter", ID: "A674232"})

	router.HandleFunc("/cats", getCats).Methods("Get")           //read all cats
	router.HandleFunc("/cats/{id}", getCat).Methods("GET")       // read one cat
	router.HandleFunc("/cats/{id}", updateCat).Methods("PUT")    // update cat info
	router.HandleFunc("/cats/{id}", deleteCat).Methods("DELETE") // delete cat
	router.HandleFunc("/cats", createCat).Methods("POST")        //create a new cat

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
