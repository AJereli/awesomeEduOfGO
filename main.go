package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
	fmt.Println(r.Body)
}

func main() {
	h := http.NewServeMux()
	h.HandleFunc("/testRec", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Test answer from testRec")
	})
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "unrec request")
	})
	err := http.ListenAndServe(":9999", h)

	log.Fatal(err)
}