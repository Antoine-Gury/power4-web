// Package contenant la logique du jeu
package src

// Structure du jeu
type GameMedium struct {
	Board         [6][9]int // Grille : 0=vide, 1=rouge, 2=jaune
	CurrentPlayer int       // Joueur actuel (1 ou 2)
	Winner        int       // 0=aucun, 1 ou 2=gagnant, -1=nul
	GameOver      bool      // true si partie finie
}

// Variable globale pour stocker le jeu
var currentGamemedium *GameMedium

// Crée une nouvelle partie
func InitGamemedium() {
	currentGamemedium = &GameMedium{
		CurrentPlayer: 1,
		Winner:        0,
		GameOver:      false,
	}
}

// Récupère le jeu actuel
func GetGamemedium() *GameMedium {
	return currentGamemedium
}

// Joue dans une colonne (0 à 6)
func (g *GameMedium) PlayColumn(col int) bool {
	// Si partie finie ou colonne invalide
	if g.GameOver || col < 0 || col >= 9 {
		return false
	}

	// Trouve la première case vide en partant du bas
	for row := 5; row >= 0; row-- {
		if g.Board[row][col] == 0 {
			// Place le pion
			g.Board[row][col] = g.CurrentPlayer

			// Vérifie victoire
			if g.checkWin(row, col) {
				g.Winner = g.CurrentPlayer
				g.GameOver = true
				return true
			}

			// Vérifie match nul
			if g.isFull() {
				g.Winner = -1
				g.GameOver = true
				return true
			}

			// Change de joueur
			if g.CurrentPlayer == 1 {
				g.CurrentPlayer = 2
			} else {
				g.CurrentPlayer = 1
			}
			return true
		}
	}
	return false // Colonne pleine
}

// Vérifie si quelqu'un a gagné
func (g *GameMedium) checkWin(row, col int) bool {
	player := g.Board[row][col]

	// Vérifie horizontal
	count := 1
	// Compte à gauche
	for c := col - 1; c >= 0 && g.Board[row][c] == player; c-- {
		count++
	}
	// Compte à droite
	for c := col + 1; c < 9 && g.Board[row][c] == player; c++ {
		count++
	}
	if count >= 4 {
		return true
	}

	// Vérifie vertical
	count = 1
	// Compte en bas
	for r := row + 1; r < 6 && g.Board[r][col] == player; r++ {
		count++
	}
	if count >= 4 {
		return true
	}

	// Vérifie diagonal \
	count = 1
	for i := 1; i < 4; i++ {
		r, c := row-i, col-i
		if r < 0 || c < 0 || g.Board[r][c] != player {
			break
		}
		count++
	}
	for i := 1; i < 4; i++ {
		r, c := row+i, col+i
		if r >= 6 || c >= 9 || g.Board[r][c] != player {
			break
		}
		count++
	}
	if count >= 4 {
		return true
	}

	// Vérifie diagonal /
	count = 1
	for i := 1; i < 4; i++ {
		r, c := row-i, col+i
		if r < 0 || c >= 9 || g.Board[r][c] != player {
			break
		}
		count++
	}
	for i := 1; i < 4; i++ {
		r, c := row+i, col-i
		if r >= 6 || c < 0 || g.Board[r][c] != player {
			break
		}
		count++
	}
	return count >= 4
}

// Vérifie si la grille est pleine
func (g *GameMedium) isFull() bool {
	for col := 0; col < 9; col++ {
		if g.Board[0][col] == 0 {
			return false
		}
	}
	return true
}
