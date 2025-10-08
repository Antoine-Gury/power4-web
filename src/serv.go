// Package contenant la logique du serveur
package src

// Import des packages nécessaires
import (
	"fmt"      // Pour formater et afficher du texte
	"log"      // Pour gérer les logs et erreurs
	"net"      // Pour trouver un port libre
	"net/http" // Pour créer le serveur HTTP
)

// Fonction pour trouver automatiquement un port libre
func findFreePort() (int, error) {
	// Crée un listener TCP sur le port 0 (attribution automatique)
	listener, err := net.Listen("tcp", ":0")
	// Si erreur lors de la création
	if err != nil {
		// Retourne 0 et l'erreur
		return 0, err
	}
	// Ferme le listener après récupération du port
	defer listener.Close()

	// Extrait le numéro de port attribué
	port := listener.Addr().(*net.TCPAddr).Port
	// Retourne le port trouvé
	return port, nil
}

// Fonction principale qui démarre le serveur et configure les routes
func StartServer() {
	// Initialise une nouvelle partie de Power4
	InitGame()

	// Définit les routes HTTP
	http.HandleFunc("/", HomeHandler)       // Page d'accueil
	http.HandleFunc("/play", PlayHandler)   // Jouer un coup
	http.HandleFunc("/reset", ResetHandler) // Nouvelle partie

	// Trouve un port libre
	port, err := findFreePort()
	// Si erreur lors de la recherche
	if err != nil {
		// Arrête le programme avec l'erreur
		log.Fatal("❌ Erreur lors de la recherche d'un port libre:", err)
	}

	// Construit l'URL du jeu
	gameURL := fmt.Sprintf("http://localhost:%d", port)
	// Affiche les informations de démarrage
	fmt.Println("🎮 Serveur Power4 démarré avec succès !")
	fmt.Printf("🔗 Accédez au jeu ici : %s\n", gameURL)
	fmt.Println("⏹️  Appuyez sur Ctrl+C pour arrêter le serveur")

	// Construit l'adresse du serveur
	address := fmt.Sprintf(":%d", port)
	// Lance le serveur HTTP (bloque jusqu'à l'arrêt)
	log.Fatal(http.ListenAndServe(address, nil))
}
