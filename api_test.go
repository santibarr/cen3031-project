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

	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/cats", getCats)

	router.ServeHTTP(responseRecorder, req)

}
