package main

import (
	"SC"
	"bufio"
	"flag"
	"fmt"

	"net/http"
	"os"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
	fmt.Println(r.Body)
}

func console(server *SC.Server, httpSer *http.Server) {

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if text == "exit\n" {
			server.Close()
			httpSer.Shutdown(nil)
			return
		}
	}
}

var webroot = flag.String("root", os.Getenv("HOME"), "web root directory")

//func main() {
//	//h := http.NewServeMux()
//	//h.HandleFunc("/testRec", func(w http.ResponseWriter, r *http.Request) {
//	//	fmt.Fprintln(w, "Test answer from testRec")
//	//})
//	//h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//	//	w.WriteHeader(404)
//	//	fmt.Fprintln(w, "unrec request")
//	//})
//	//err := http.ListenAndServe(":9999", h)
//
//	s := SC.MakeServer("/ent")
//	srv := &http.Server{Addr: ":8080"}
//	go s.Listen()
//	go console(s, srv)
//	// static files
//
//
//	http.Handle("/", http.FileServer(http.Dir(*webroot)))
//	srv.ListenAndServe()
//
//}
