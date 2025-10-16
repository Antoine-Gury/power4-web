// Point d'entrée du programme
package main

// Import du package qui gère le serveur
import (
	src "power4/src/go" // Package contenant la logique du serveur
)

// Fonction principale - lance simplement le serveur
func main() {
	// Appelle la fonction qui démarre tout
	src.StartServer()
}
