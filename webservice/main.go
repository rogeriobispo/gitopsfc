package main

import "net/http"

func main() {
	println("Listening port 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello world</h1>"))
	})

	http.ListenAndServe(":8080", nil)
}