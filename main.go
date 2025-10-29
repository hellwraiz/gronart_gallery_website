package main

import (
	"fmt"
	"net/http"
	"os"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	
    port := os.Getenv("PORT")
	fmt.Println("Attempting to start the server on port", port)
    if port == "" {
        port = "8080"
    }


	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on port", port)
	http.ListenAndServe(":8080", nil)
}
