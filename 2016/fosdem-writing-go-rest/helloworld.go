// +build OMIT

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	world := "World"
	http.HandleFunc("/helloworld",
		func(w http.ResponseWriter, r *http.Request) { // HL
			fmt.Fprintf(w, "Hello %s", world) // HL
		}) // HL

	log.Fatal(http.ListenAndServe(":8080", nil))
}
