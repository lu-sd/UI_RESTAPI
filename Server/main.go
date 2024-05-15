package main

import (
	"fmt"
	"net/http"
	"time"
)

var fileHandler = http.FileServer(http.Dir("../UI/dist"))

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+
	// 	"<textarea name=\"body\">%s</textarea><br>"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// 	"</form>",
	// 	"page", "page", "content")
	fmt.Fprint(w, time.Now().Format("02 Jan 2006 15:04:05 "))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", fileHandler)

	mux.HandleFunc("GET /hello", helloHandler)

	mux.HandleFunc("/task/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})

	http.ListenAndServe("localhost:8090", mux)
}
