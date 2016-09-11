package main

import (
	"fmt"
	"io"
	"net/http"
)

var (
	count = 0
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	count++
	fmt.Println(req.RequestURI, req.URL)
	s := fmt.Sprintf("hello %d  %s\n", count, req.RequestURI)
	io.WriteString(w, s)
}

func main() {
	err := http.ListenAndServe(":80", http.HandlerFunc(HelloServer))
	fmt.Print(err)
}
