package main

import (
	"fmt"
	"net/http"
)

// Handler pour la page d'accueil
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenue sur la page d’accueil !")
}

// Handler pour la page "À propos"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ceci est la page À propos.")
}

// Handler pour la page "Contact"
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page Contact : envoyez-nous un message !")
}

func main() {
	// Association de la route "/" avec homeHandler
	http.HandleFunc("/", homeHandler)
	// Association de la route "/about" avec aboutHandler
	http.HandleFunc("/about", aboutHandler)
	// Association de la route "/contact" avec contactHandler
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Serveur démarré sur le port 8080...")
	// Démarre l'écoute sur le port 8080
	http.ListenAndServe(":8080", nil)
}
