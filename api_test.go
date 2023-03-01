package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func testgetCats(t *testing.T) {

	req, err := http.NewRequest("GET", "/cats", nil)

	//error checking
	if err != nil {

		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/cats", getCats).Methods("GET")

	router.ServeHTTP(response, req)

	status := response.Code

	if status != http.StatusOK {

		t.Errorf("Status does not match up to StatusOK (200)")
	}

	println(response.Body)

}
