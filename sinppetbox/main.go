package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello "))
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Display a specific snippet"))
	fmt.Fprintf(w, "Display a specific with id %d ", id)
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//w.Header().Set("Allow", "POST")
		w.Header().Set("Allow", http.MethodPost)
		//w.WriteHeader(405)
		//w.Write([]byte("method not allowed"))
		//http.Error(w, "Method not allowed ", 405)
		http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("create a new snippet"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("Start server on : 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatalln(err)
}
