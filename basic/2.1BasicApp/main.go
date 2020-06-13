package main

import (
	"fmt"
	"net/http"
)

// this will listen the webapp
func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //specific to we request
	//!! Browser renders the page as an html page.
	fmt.Fprint(w, r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprint(w, "Hey this is golang")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "Click here <a href = \"https://github.com/rajneesh44\"> abcdd </a>")
	}
	// r is the request we are sending
	// w is the retrun response we get from the server
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil) // nil is the builtin serve mux

}

// handler func takes repsone writer and request
// w interfaces to the io

// dyanmic reloading : automatically chaanges and reloading (not default)
// downloaded "fresh package" Simply run the app using 1. go into the main directory where your files are there
// $ fresh
// fresh cant detect compiation error, so try to start server using go run file.go
