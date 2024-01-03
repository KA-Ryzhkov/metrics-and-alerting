package main

import (
	"fmt"
	"net/http"
)

func updatePage(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, updatePage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
