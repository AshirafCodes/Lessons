package main

import (
	"fmt"
	"log"
	"net/http"
)

func output(path string) string {
	return "This path is " + path
}

func main() {
	fmt.Println("Start API")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Bulungi")
		w.WriteHeader(200)
		w.Write([]byte(output("/")))
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Bulungi")
		w.WriteHeader(200)
		w.Write([]byte(output("/home")))
	})
	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Bulungi")
		w.WriteHeader(200)
		w.Write([]byte(output("/galler")))
	})
	http.HandleFunc("/newsletter", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Bulungi")
		w.WriteHeader(200)
		w.Write([]byte(output("/newsletter")))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
