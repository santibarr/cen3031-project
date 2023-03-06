package api

//import packages - update
import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

// Needs to match with cat struct in scrape.go
type Cat struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    string `json:"age"`
	Loc    string `json:"location"`
	ID     string `json:"id"`
}

// hold a slice of cats, will maybe change in future
var cats []Cat

// get all the cats
func GetCats(w http.ResponseWriter, r *http.Request) {

	//Status -> 200
	w.WriteHeader(http.StatusOK)

	//we want to set the header to be a json type
	w.Header().Set("Content-Type", "application/json")

	//encode the cats slice and put it on the front end
	json.NewEncoder(w).Encode(cats)
}

// delete a cat from the website
func deleteCat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusNoContent)

	//fill in the parameters from the website
	params := mux.Vars(r)

	for index, item := range cats {

		//if the name and the location of the cat are the same then get rid of it
		if item.ID == params["id"] {

			fmt.Println("In Delete if statement")
			cats = append(cats[:index], cats[index+1:]...)
			break
		}

	}
	//encode the new cats slice
	json.NewEncoder(w).Encode(cats)

}

// if we want a singular cat
func getCat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//Status -> 200
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)

	for _, item := range cats {

		if item.ID == params["id"] {

			//go get the cat in question
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

// create a new cat object
func createCat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//Status -> 200
	w.WriteHeader(http.StatusOK)

	var cat Cat //variable cat

	//decode the information from the json to write it back
	_ = json.NewDecoder(r.Body).Decode(&cat)

	cats = append(cats, cat)

	json.NewEncoder(w).Encode(cat)

}

func updateCat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//Status -> 200
	w.WriteHeader(http.StatusOK)

	//inside the for loop we delete than update
	for index, item := range cats {

		if item.ID == params["id"] {

			//delete the cat
			cats = append(cats[:index], cats[index+1:]...)

			var cat Cat

			_ = json.NewDecoder(r.Body).Decode(&cat)

			cat.ID = params["id"]

			//update with new cat
			cats = append(cats, cat)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
