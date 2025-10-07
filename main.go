// Déclaration du package principal (point d'entrée du programme)
package main

// Import des packages nécessaires
import (
	"fmt"        // Package pour formater et afficher du texte
	"log"        // Package pour gérer les logs et erreurs fatales
	"net"        // Package pour les opérations réseau (trouver un port)
	"net/http"   // Package pour créer le serveur HTTP
	"power4/src" // Import de notre package personnalisé contenant les handlers
)

// Fonction pour trouver automatiquement un port libre sur la machine
func findFreePort() (int, error) {
	// Crée un listener TCP sur le port 0 (le système attribue automatiquement un port libre)
	listener, err := net.Listen("tcp", ":0")
	// Si une erreur survient lors de la création du listener
	if err != nil {
		// Retourne 0 et l'erreur
		return 0, err
	}
	// Ferme le listener après avoir récupéré le port (defer = exécuté à la fin de la fonction)
	defer listener.Close()

	// Récupère l'adresse du listener et extrait le numéro de port
	port := listener.Addr().(*net.TCPAddr).Port
	// Retourne le port trouvé et nil (pas d'erreur)
	return port, nil
}

// Fonction principale qui démarre le serveur et le jeu
func main() {
	// Initialise une nouvelle partie de Power4
	src.InitGame()

	// Définit la route GET pour afficher la page d'accueil du jeu
	http.HandleFunc("/", src.HomeHandler)
	// Définit la route POST pour jouer un coup (recevoir la colonne choisie)
	http.HandleFunc("/play", src.PlayHandler)
	// Définit la route POST pour réinitialiser le jeu (nouvelle partie)
	http.HandleFunc("/reset", src.ResetHandler)

	// Appelle la fonction pour trouver un port libre
	port, err := findFreePort()
	// Si une erreur survient lors de la recherche du port
	if err != nil {
		// Arrête le programme et affiche l'erreur
		log.Fatal("❌ Erreur lors de la recherche d'un port libre:", err)
	}

	// Construit l'URL complète du jeu avec le port trouvé
	gameURL := fmt.Sprintf("http://localhost:%d", port)
	// Affiche un message de démarrage réussi
	fmt.Println("🎮 Serveur Power4 démarré avec succès !")
	// Affiche le lien cliquable pour accéder au jeu
	fmt.Printf("🔗 Accédez au jeu ici : %s\n", gameURL)
	// Affiche les instructions pour arrêter le serveur
	fmt.Println("⏹️  Appuyez sur Ctrl+C pour arrêter le serveur")

	// Construit l'adresse au format ":port" pour le serveur
	address := fmt.Sprintf(":%d", port)
	// Lance le serveur HTTP sur le port trouvé (bloque ici jusqu'à l'arrêt du serveur)
	log.Fatal(http.ListenAndServe(address, nil))
}
