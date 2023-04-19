// package api

// //import packages - update
// import (
// 	"encoding/json"
// 	"fmt"

// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/santibarr/cen3031-project/scrape"
// )

// // hold a slice of cats, will maybe change in future
// var cats []scrape.Cat

// // create a new user object for contacts page
// type CatCharacteristic struct {
// 	Feature string `json:"Feature"`
// }

<<<<<<< HEAD
// type User struct {
// 	FirstName   string `json:"firstName"`
// 	LastName    string `json:"lastName"`
// 	Email       string `json:"email"`
// 	phoneNumber string `json:"phone"`
// }
=======
type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone"`
}
>>>>>>> d9986b1a3fafe3f232bce0af1d4df3f7365cc2d2

// var characteristic []CatCharacteristic
// var users []User //user slice

// // create cat quiz struct
// // get all the cats
// func GetCats(w http.ResponseWriter, r *http.Request) {

// 	//Status -> 200s
// 	w.WriteHeader(http.StatusOK)

// 	cats = append(cats, scrape.MiamiDade(cats)...)
// 	cats = append(cats, scrape.LakeCounty(cats)...)
// 	cats = append(cats, scrape.Peggy(cats)...)
// 	cats = append(cats, scrape.Marathon(cats)...)
// 	cats = append(cats, scrape.KeyWest(cats)...)

// 	//we want to set the header to be a json type
// 	w.Header().Set("Content-Type", "application/json")

// 	//encode the cats slice and put it on the front end
// 	json.NewEncoder(w).Encode(cats)

// }

// // delete a cat from the website
// func DeleteCat(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	w.WriteHeader(http.StatusNoContent)

// 	cats = append(cats, scrape.MiamiDade(cats)...)
// 	cats = append(cats, scrape.LakeCounty(cats)...)
// 	cats = append(cats, scrape.Peggy(cats)...)
// 	cats = append(cats, scrape.Marathon(cats)...)
// 	cats = append(cats, scrape.KeyWest(cats)...)

// 	//fill in the parameters from the website
// 	params := mux.Vars(r)

// 	for index, item := range cats {

// 		//if the name and the location of the cat are the same then get rid of it
// 		if item.Name == params["Name"] {

// 			fmt.Println("In Delete if statement")
// 			cats = append(cats[:index], cats[index+1:]...)
// 			break
// 		}

// 	}
// 	//encode the new cats slice
// 	json.NewEncoder(w).Encode(cats)

// }

// // if we want a singular cat
// func GetCat(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	//Status -> 200
// 	w.WriteHeader(http.StatusOK)
// 	cats = append(cats, scrape.MiamiDade(cats)...)
// 	cats = append(cats, scrape.LakeCounty(cats)...)
// 	cats = append(cats, scrape.Peggy(cats)...)
// 	cats = append(cats, scrape.Marathon(cats)...)
// 	cats = append(cats, scrape.KeyWest(cats)...)

// 	params := mux.Vars(r)

// 	for _, item := range cats {

// 		if item.Name == params["Name"] {

// 			//go get the cat in question
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}

// }
// func GetCatQuiz(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	w.WriteHeader(http.StatusOK)
// 	cats = append(cats, scrape.MiamiDade(cats)...)
// 	cats = append(cats, scrape.LakeCounty(cats)...)
// 	cats = append(cats, scrape.Peggy(cats)...)
// 	cats = append(cats, scrape.Marathon(cats)...)
// 	cats = append(cats, scrape.KeyWest(cats)...)

// 	// var userResponse []CatCharacteristic
// 	// json.NewDecoder(r.Body).Decode(&userResponse)
// 	// fmt.Print(userResponse)

// 	// params := mux.Vars(r)

// 	// selectedFeature := params["Feature"]
// 	// var newCats []scrape.Cat
// 	// for _, item := range cats {

// 	// 	if item.Feature == selectedFeature {
// 	// 		newCats = append(newCats, item)
// 	// 	}
// 	// }
// 	params := mux.Vars(r)
// 	selectedFeature := params["Feature"]

// 	// Filter cats based on selected feature
// 	var newCats []scrape.Cat
// 	for _, item := range cats {
// 		if item.Feature == selectedFeature {
// 			newCats = append(newCats, item)
// 		}
// 	}


// 	json.NewEncoder(w).Encode(newCats)

// }
// func GetCatQuizAddition(w http.ResponseWriter, r *http.Request) {

// 	cats = append(cats, scrape.MiamiDade(cats)...)
// 	cats = append(cats, scrape.LakeCounty(cats)...)
// 	cats = append(cats, scrape.Peggy(cats)...)
// 	cats = append(cats, scrape.Marathon(cats)...)
// 	cats = append(cats, scrape.KeyWest(cats)...)

// 	selectedFeature := r.URL.Query().Get("feature")

// 	pickedCats := make([]scrape.Cat, 0)
// 	for _, cats := range cats {
// 		if cats.Feature == selectedFeature {
// 			pickedCats = append(pickedCats, cats)
// 		}
// 	}

// 	responseJSON, err := json.Marshal(pickedCats)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(responseJSON)
// }

// // create a new cat object
// func CreateUser(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	var user User //variable user

// 	//decode the information from the json to write it back
// 	err := json.NewDecoder(r.Body).Decode(&user)

// 	//print the error if there is one
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 	}

// 	users = append(users, user) //append the user to the slice

// 	//Status -> 200
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(user)

// 	//print the user
// 	for i := range users {
// 		fmt.Println(users[i])
// 	}

// }
// func CreateCat(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	//Status -> 200
// 	w.WriteHeader(http.StatusOK)
// 	cats = append(cats, scrape.MiamiDade(cats)...)
// 	cats = append(cats, scrape.LakeCounty(cats)...)
// 	cats = append(cats, scrape.Peggy(cats)...)
// 	cats = append(cats, scrape.Marathon(cats)...)
// 	cats = append(cats, scrape.KeyWest(cats)...)

// 	var cat scrape.Cat //variable cat

// 	//decode the information from the json to write it back
// 	_ = json.NewDecoder(r.Body).Decode(&cat)

// 	cats = append(cats, cat)

// 	json.NewEncoder(w).Encode(cat)

// }

// func UpdateCat(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	//Status -> 200
// 	w.WriteHeader(http.StatusOK)

// 	cats = append(cats, scrape.MiamiDade(cats)...)
// 	cats = append(cats, scrape.LakeCounty(cats)...)
// 	cats = append(cats, scrape.Peggy(cats)...)
// 	cats = append(cats, scrape.Marathon(cats)...)
// 	cats = append(cats, scrape.KeyWest(cats)...)

// 	//inside the for loop we delete than update
// 	for index, item := range cats {

// 		if item.Name == params["name"] {

// 			//delete the cat
// 			cats = append(cats[:index], cats[index+1:]...)

// 			var cat scrape.Cat

// 			_ = json.NewDecoder(r.Body).Decode(&cat)

// 			cat.Name = params["name"]

// 			//update with new cat
// 			cats = append(cats, cat)
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}

// }
