package main

import (
	"fmt"
	"html/template"
	"museum/go/server/api"
	"museum/go/server/data"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) { 
	w.Write([]byte("Hello!"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request ) { 
	html, err := template.ParseFiles("templates/index.tmpl")

	if err != nil { 
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)


	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":8888", server)

	if err != nil { 
		fmt.Println(("Error while opening the server"))
	}
}