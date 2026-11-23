package main

import (
	"fmt"
	"log"
	"net/http"
	// "initial/controllers"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "templates/mayor.html")
	})

	var id string

	log.Printf("aaaa")
	http.HandleFunc("/Player", func(w http.ResponseWriter, r *http.Request) {
		id = r.URL.Query().Get("PlayerId")

		fmt.Fprintf(w, "PlayerUserName reçu: %s\n", id)

		if id[0] >= '0' && id[0] <= '9' {

		} else if id[0] >= 'a' && id[0] <= 'z' {

		} else {
			fmt.Fprintf(w, "erreur\n")
		}

		fmt.Fprintf(w, "Le pokemon est : %+v\n", pokemon)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
