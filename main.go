package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started and listenning on localhost:3000")
	log.Fatal(http.ListenAndServe(":9003", nil))

}
