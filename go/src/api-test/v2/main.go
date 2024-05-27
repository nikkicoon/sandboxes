package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// struct server has no fields,
// type server struct{}

// add a method to this server called 'ServeHTTP' to satisfy the Handler
// Interface.
// Interface looks like this:
//
// type Handler interface {
//        ServeHTTP(ResponseWriter, *Request)
// }
// It has one method and one method only.
// A struct or object will be Handler if it has one
// method ServeHTTP which takes ResponseWriter and pointer to Request.
//
// We don't have to say that we are implementing the interface.
// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
func home (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST called"})`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "PUT called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "DELETE called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	//s := &server{}
	//http.Handle("/", s)
	// http.HandleFunc("/", home)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":8080", r))
}
