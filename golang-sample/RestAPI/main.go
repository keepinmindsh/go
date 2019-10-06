package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Post struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

var posts []Post

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/gets", getPosts ).Methods("GET")

	_ = http.ListenAndServe(":8080", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	fmt.Println(params)

	json.NewEncoder(w).Encode(params["field"])
}