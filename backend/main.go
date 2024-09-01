package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Web Services are easy with Go!")
	})

	http.ListenAndServe(":8000", nil)
}