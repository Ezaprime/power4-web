package power4

import (
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/reset", resetHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Serveur en ligne sur http://localhost:8080")
	mustLogPaths()
	log.Fatal(http.ListenAndServe(":8080", nil))
}