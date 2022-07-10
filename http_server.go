package main

import (
"fmt"
"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Accepting an incoming hello request")
	fmt.Fprintf(w, "hello\n")
  fmt.Println("Finished an incoming hello request")
}

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}
