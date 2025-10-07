// Package contenant les handlers HTTP
package src

// Import des packages nécessaires
import (
	"html/template" // Pour charger et afficher les templates HTML
	"net/http"      // Pour gérer les requêtes HTTP
	"strconv"       // Pour convertir string en int
)

// Handler pour afficher la page du jeu (route GET /)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Charge le template HTML
	tmpl, err := template.ParseFiles("templates/index.html")
	// Si erreur lors du chargement du template
	if err != nil {
		// Affiche l'erreur dans le navigateur
		http.Error(w, "Erreur template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Récupère l'état actuel du jeu
	game := GetGame()

	// Affiche le template avec les données du jeu
	tmpl.Execute(w, game)
}

// Handler pour jouer un coup (route POST /play)
func PlayHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère le numéro de colonne envoyé par le formulaire
	colStr := r.FormValue("column")

	// Convertit la string en nombre
	col, err := strconv.Atoi(colStr)
	// Si erreur de conversion
	if err != nil {
		// Redirige vers la page d'accueil
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Récupère le jeu actuel
	game := GetGame()

	// Joue dans la colonne choisie
	game.PlayColumn(col)

	// Redirige vers la page d'accueil pour afficher le résultat
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Handler pour recommencer une partie (route POST /reset)
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	// Crée une nouvelle partie
	InitGame()

	// Redirige vers la page d'accueil
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
