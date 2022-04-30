package main

import "github.com/Uzama/most_used_words/http"

func main() {
	server := http.NewServer(8080)

	// create a new server
	server.Create()

	// run the server
	server.Start()
}
