package main

//import packages
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	cats = append(cats, Cat{Name: "Smithy", Age: "1 year old", Gender: "Male", Loc: "Gainesville Cat Cafe"})
	cats = append(cats, Cat{Name: "Ruby", Age: " 4 months old", Gender: "Female", Loc: "Miami Animal Shelter"})

	router.HandleFunc("/cats", getCats).Methods("Get")                   //read all cats
	router.HandleFunc("/cats/{name}/{loc}", getCat).Methods("GET")       // read one cat
	router.HandleFunc("cats/{name}/{loc}", updateCat).Methods("PUT")     // update cat info
	router.HandleFunc("/cats/{name}/{loc}", deleteCat).Methods("DELETE") // delete cat
	router.HandleFunc("/cats", createCat).Methods("POST")                //create a new cat

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
