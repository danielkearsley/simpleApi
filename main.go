package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

const port = ":3030"

type Page struct {
	Title string
	Body  string
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Index)
	router.HandleFunc("/get", Get).Methods("GET")
	router.HandleFunc("/post", Post).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedHeaders([]string{"*"})
	methodsOk := handlers.AllowedHeaders([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server on" + port)
	http.ListenAndServe(port, handlers.CORS(headersOk, originsOk, methodsOk)(router))
}

//Routes
func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	p := Page{"My Api", "Hello World"}

	t.Execute(w, p)
}

func Get(w http.ResponseWriter, r *http.Request) {
	//Get paramaters
	vars := r.URL.Query()
	//Getting data
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func Post(w http.ResponseWriter, r *http.Request) {
	//Post parameters
	message := r.FormValue("msg")

	json.NewEncoder(w).Encode(map[string]string{"Message": message})
}
