package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/santibarr/cen3031-project/api"
	"github.com/santibarr/cen3031-project/scrape"
)

func main() {

	//router created

	//slice of cats, will change in future
	var cats []scrape.Cat

	cats = scrape.MiamiDade(cats)
	cats = scrape.LakeCounty(cats)
	cats = scrape.Peggy(cats)
	cats = scrape.Marathon(cats)
	cats = scrape.KeyWest(cats)

	router := mux.NewRouter()

	//Err would be 404 not found, Directory does not exist
	fs := http.FileServer(http.Dir("dist/purfect-partner"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	//example for now, but should invoke scraper.go to get cat objects.
	// cats = append(cats, Cat{Name: "Smithy", Age: "1 year old", Gender: "Male", Loc: "Gainesville Cat Cafe", ID: "A245676"})
	// cats = append(cats, Cat{Name: "Ruby", Age: " 4 months old", Gender: "Female", Loc: "Miami Animal Shelter", ID: "A674232"})

	router.HandleFunc("/cats", api.GetCats).Methods("GET")           //read all cats
	router.HandleFunc("/cats/{id}", api.GetCat).Methods("GET")       // read one cat
	router.HandleFunc("/cats/{id}", api.UpdateCat).Methods("PUT")    // update cat info
	router.HandleFunc("/cats/{id}", api.DeleteCat).Methods("DELETE") // delete cat
	router.HandleFunc("/cats", api.CreateCat).Methods("POST")        //create a new cat

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
