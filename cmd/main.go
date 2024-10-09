package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Benvingut al MetaContext API!")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Servidor HTTP escoltant al port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
