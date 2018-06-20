package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
        fmt.Fprintf(w, "Hello world.")
	fmt.Println("helloHandler is called.")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	http.ListenAndServe("0.0.0.0:8080", mux)
}
