package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var count int

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", IncrementRequest)
	http.Handle("/", router)

	var err error
	http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatalln("Error start server")
	}

}

func IncrementRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>HI</h1>")
	count += 1
	fmt.Printf(" ⛔ \t Request received -> [%d] 👿 \n", count)

}
