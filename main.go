package main

import (
	"net/http"
)

func main() {
	print("Init: Health check server up and running")
	http.HandleFunc("/health", func(writer http.ResponseWriter, reader *http.Request) {
		writer.Header().Set("Server", "Alive and Kicking!")
		writer.WriteHeader(200)
	})

	http.ListenAndServe(":5454", nil)
}
