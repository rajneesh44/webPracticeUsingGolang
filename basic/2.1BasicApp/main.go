package main

import (
	"fmt"
	"net/http"
)

// this will listen the webapp
func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Welome to my awesome site</h1>")
}

func main(){
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)	  // nil is the builtin serve mux

}
