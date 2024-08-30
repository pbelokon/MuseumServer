package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { 
		w.Write([]byte("Hello!"))
	})
	err := http.ListenAndServe(":8888", nil)

	if err != nil { 
		fmt.Println(("Error while opening the server"))
	}
}