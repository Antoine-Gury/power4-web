package main

import (
	"html/template"
	"net/http"
)

// Handler pour la page d'accueil
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := map[string]string{
		"Title":   "Bienvenue sur Puissance 4 🎮",
		"Message": "Ceci est la page d’accueil générée avec un template Go !",
	}
	tmpl.Execute(w, data)
}

// Fonction main pour démarrer le serveur
func main() {
	http.HandleFunc("/", homeHandler)

	println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
