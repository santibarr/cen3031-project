// package api

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gorilla/mux"
// )

// // gets singular cat
// func TestGetCats(t *testing.T) {

// 	req, err := http.NewRequest("GET", "/cats", nil)

// 	//error checking
// 	if err != nil {

// 		t.Fatal(err)
// 	}

// 	response := httptest.NewRecorder()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/cats", GetCats).Methods("GET")

// 	router.ServeHTTP(response, req)

// 	status := response.Code

// 	println("This is the status code ", status)

// 	if status != http.StatusOK {

// 		t.Errorf("Status does not match up to StatusOK (200)")
// 	}

// 	println(response.Body.String())

// }

// func TestGetCat(t *testing.T) {

// 	req, err := http.NewRequest("GET", "/cats/A245676", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	response := httptest.NewRecorder()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/cats/{id}", GetCat).Methods("GET")

// 	router.ServeHTTP(response, req)

// 	status := response.Code

// 	println("This is the status code ", status)

// 	if status != http.StatusOK {

// 		t.Errorf("Status does not match up to StatusOK (200)")
// 	}

// 	println(response.Body.String())

// }

// func TestCreateCat(t *testing.T) {

// 	requestBody := []byte(`{"name": "John Doe", "gender": "Female", "id": "A542786", "age": "3 months", "location": "Miami Dade"}`)
// 	req, err := http.NewRequest("POST", "/cats", bytes.NewBuffer(requestBody))

// 	if err != nil {

// 		t.Fatal(err)
// 	}

// 	response := httptest.NewRecorder() //hold the response

// 	//create the new router
// 	router := mux.NewRouter()
// 	router.HandleFunc("/cats", CreateCat).Methods("POST")

// 	router.ServeHTTP(response, req) //create the http and see what happens

// 	status := response.Code //get the status code

// 	println("This is the status code ", status)

// 	if status != http.StatusOK {

// 		t.Errorf("Status does not match up to StatusOK (200)")
// 	}

// 	println(response.Body.String())

// }

// // test the update api endpoint
// func TestUpdateCat(t *testing.T) {

// 	requestBody := []byte(`{"name": "John Doe", "gender": "Female", "id": "A542786", "age": "3 months", "location": "Miami Dade"}`)
// 	req, err := http.NewRequest("PUT", "/cats/A245676", bytes.NewBuffer(requestBody))

// 	if err != nil {

// 		t.Fatal(err)
// 	}

// 	response := httptest.NewRecorder() //hold the response

// 	//create the new router
// 	router := mux.NewRouter()
// 	router.HandleFunc("/cats/{id}", UpdateCat).Methods("PUT")

// 	router.ServeHTTP(response, req) //create the http and see what happens

// 	status := response.Code //get the status code

// 	println("This is the status code ", status)

// 	if status != http.StatusOK {

// 		t.Errorf("Status does not match up to StatusOK (200)")
// 	}

// 	println(response.Body.String())

// }

// // test the delete api endpoint
// func TestDeleteCat(t *testing.T) {

// 	req, err := http.NewRequest("DELETE", "/cats/A245676", nil)

// 	if err != nil {

// 		t.Fatal(err)
// 	}

// 	response := httptest.NewRecorder()
// 	router := mux.NewRouter()
// 	router.HandleFunc("/cats/{id}", DeleteCat).Methods("DELETE")
// 	router.ServeHTTP(response, req)

// 	status := response.Code

// 	if status != http.StatusNoContent {

// 		t.Errorf("Status does not match up to StatusNoContent")

// 	}

// }
// //-----New Unit test for Sprint 4-------//
// func TestGetCatQuiz(t *testing.T){
// 	req, err := http.NewRequest("GET", "/cats/A245676", nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	response := httptest.NewRecorder()

// 	router := mux.NewRouter()
// 	router.HandleFunc("/cats/{id}", GetCatQuiz).Methods("GET")

// 	router.ServeHTTP(response, req)

// 	status := response.Code

// 	println("This is the status code ", status)

// 	if status != http.StatusOK {

// 		t.Errorf("Status does not match up to StatusOK (200)")
// 	}

// 	println(response.Body.String())

// }