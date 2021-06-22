package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "", http.StatusBadRequest)
		}
		fmt.Printf("Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello goodbye")
	})

	http.ListenAndServe(":8080", nil)
}
