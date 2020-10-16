package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var id int

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, NUS Hackers! (%d) Kubernetes is awesome!", id)
}

func main() {
	// generate a random id
	rand.Seed(time.Now().UnixNano())
	id = rand.Intn(9999)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
