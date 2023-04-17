package api

//import packages - update
import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/santibarr/cen3031-project/scrape"
)

// hold a slice of cats, will maybe change in future
var cats []scrape.Cat

// create a new user object for contacts page
type User struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Email     string `json:"Email"`
	Phone     string `json:"Phone"`
}

var users []User //user slice

// get all the cats
func GetCats(w http.ResponseWriter, r *http.Request) {

	//Status -> 200s
	w.WriteHeader(http.StatusOK)

	cats = append(cats, scrape.MiamiDade(cats)...)
	cats = append(cats, scrape.LakeCounty(cats)...)
	cats = append(cats, scrape.Peggy(cats)...)
	cats = append(cats, scrape.Marathon(cats)...)
	cats = append(cats, scrape.KeyWest(cats)...)

	//we want to set the header to be a json type
	w.Header().Set("Content-Type", "application/json")

	//encode the cats slice and put it on the front end
	json.NewEncoder(w).Encode(cats)

}

// delete a cat from the website
func DeleteCat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusNoContent)

	cats = append(cats, scrape.MiamiDade(cats)...)
	cats = append(cats, scrape.LakeCounty(cats)...)
	cats = append(cats, scrape.Peggy(cats)...)
	cats = append(cats, scrape.Marathon(cats)...)
	cats = append(cats, scrape.KeyWest(cats)...)

	//fill in the parameters from the website
	params := mux.Vars(r)

	for index, item := range cats {

		//if the name and the location of the cat are the same then get rid of it
		if item.Name == params["Name"] {

			fmt.Println("In Delete if statement")
			cats = append(cats[:index], cats[index+1:]...)
			break
		}

	}
	//encode the new cats slice
	json.NewEncoder(w).Encode(cats)

}

// we want to display the cat info quiz on the website
func GetCatQuiz(w http.ResponseWriter, r *http.Request) {

	//Status -> 200
	w.WriteHeader(http.StatusOK)
	cats = append(cats, scrape.MiamiDade(cats)...)
	cats = append(cats, scrape.LakeCounty(cats)...)
	cats = append(cats, scrape.Peggy(cats)...)
	cats = append(cats, scrape.Marathon(cats)...)
	cats = append(cats, scrape.KeyWest(cats)...)

	//we want to set the header to be a json type
	w.Header().Set("Content-Type", "application/json")

	//encode the cats slice and put it on the front end
	json.NewEncoder(w).Encode(cats)

}

// create a new cat object
func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user User //variable user

	//decode the information from the json to write it back
	err := json.NewDecoder(r.Body).Decode(&user)

	//print the error if there is one
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	users = append(users, user) //append the user to the slice

	//Status -> 200
	w.WriteHeader(http.StatusCreated)

	//print the user
	for i := range users {
		fmt.Println(users[i])
	}

}

func UpdateCat(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//Status -> 200
	w.WriteHeader(http.StatusOK)

	cats = append(cats, scrape.MiamiDade(cats)...)
	cats = append(cats, scrape.LakeCounty(cats)...)
	cats = append(cats, scrape.Peggy(cats)...)
	cats = append(cats, scrape.Marathon(cats)...)
	cats = append(cats, scrape.KeyWest(cats)...)

	//inside the for loop we delete than update
	for index, item := range cats {

		if item.Name == params["name"] {

			//delete the cat
			cats = append(cats[:index], cats[index+1:]...)

			var cat scrape.Cat

			_ = json.NewDecoder(r.Body).Decode(&cat)

			cat.Name = params["name"]

			//update with new cat
			cats = append(cats, cat)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
