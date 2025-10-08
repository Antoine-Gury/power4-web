// Package contenant les handlers HTTP
package src

// Import des packages nécessaires
import (
	"html/template" // Pour charger et afficher les templates HTML
	"net/http"      // Pour gérer les requêtes HTTP
	"strconv"       // Pour convertir string en int
)

// Variable pour stocker la difficulté actuelle
var currentDifficulty string = "classic"

// Handler pour afficher le menu principal (route GET /)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Charge le template du menu
	tmpl, err := template.ParseFiles("templates/menu.html")
	// Si erreur lors du chargement
	if err != nil {
		// Affiche l'erreur
		http.Error(w, "Erreur template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Affiche le menu
	tmpl.Execute(w, nil)
}

// Handler pour sélectionner la difficulté (route GET /play-mode)
func PlayModeHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère le paramètre "mode" de l'URL
	mode := r.URL.Query().Get("mode")

	// Définit la difficulté actuelle
	currentDifficulty = mode

	// AJOUT : initialise une nouvelle partie pour le mode choisi
	InitGame()

	// Redirige vers la page de jeu
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// Handler pour afficher la page de jeu selon la difficulté (route GET /game)
func GameHandler(w http.ResponseWriter, r *http.Request) {
	// Détermine quel template charger selon la difficulté
	var templateFile string
	switch currentDifficulty {
	case "easy":
		templateFile = "templates/gameeasy.html"
	case "medium":
		templateFile = "templates/gamemedium.html"
	case "hard":
		templateFile = "templates/gamehard.html"
	default:
		templateFile = "templates/gameclassic.html"
	}

	// Charge le template correspondant
	tmpl, err := template.ParseFiles(templateFile)
	// Si erreur lors du chargement
	if err != nil {
		// Affiche l'erreur
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
		// Redirige vers la page de jeu
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	// Récupère le jeu actuel
	game := GetGame()

	// Joue dans la colonne choisie
	game.PlayColumn(col)

	// Redirige vers la page de jeu pour afficher le résultat
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// Handler pour recommencer une partie (route POST /reset)
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	// Crée une nouvelle partie
	InitGame()

	// Redirige vers la page de jeu
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// Handler pour retourner au menu (route GET /menu)
func MenuHandler(w http.ResponseWriter, r *http.Request) {
	// Réinitialise le jeu
	InitGame()

	// Redirige vers le menu principal
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
