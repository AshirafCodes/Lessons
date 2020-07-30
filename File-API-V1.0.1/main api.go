package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start API")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Bulungi")
		w.WriteHeader(200)
		w.Write([]byte("This path is /"))
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Bulungi")
		w.WriteHeader(200)
		w.Write([]byte("This path is /home"))
	})
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Bulungi")
		w.WriteHeader(200)
		w.Write([]byte("This path is /info"))
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Bulungi")
		w.WriteHeader(200)
		w.Write([]byte("This path is /about"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

	/// The above is api S
}
