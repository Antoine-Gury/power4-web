## Contexte & objectif du projet

Ce projet est un Puissance 4 jouable sur un serveur local, développé dans le cadre scolaire.  
Il permet de lancer le jeu via un programme Go et d’y jouer directement depuis un navigateur.  
Le projet démontre l’utilisation d’un serveur HTTP en Go, la gestion d’un jeu tour par tour,  
et l’implémentation complète de la logique du Puissance 4.  
L’interface web est simple et accessible en cliquant sur le lien fourni par le terminal.  

## Prérequis

- Langages : Go v0.50.0, HTML, CSS, JavaScript  
- Dépendances : aucune (utilisation uniquement des bibliothèques standard de Go)  
- Système : Windows

- ## Installation & exécution

Cloner le projet :
```bash
git clone <URL_DU_PROJET>
```
Lancer le projet :
```bash
go run main.go
```
## Structure du dépôt

- `/src` : code source
  - `go` : fichiers `.go`
  - `js` : fichiers `.js`
- `/templates` : fichiers HTML
  - `dossier go/js` : `.html` correspondant
- `/assets` : ressources multimédias
  - `styles` : fichiers `.css`
- `/docs` : documentation du projet
  - `documentation projet` : consignes / diaporama
- `README.md` : documentation principale
- `main` : point d’entrée du programme

## Fonctions clés

### `InitGame()`
**Rôle :** Initialise une nouvelle partie de Puissance 4.  
**Signature :** `func InitGame()`  
**Détails :** Crée une structure `Game` avec une grille vide, le joueur 1 en premier, et aucune condition de victoire.  
**Complexité :** O(1)  
**Dépendances :** aucune.  
**Tests :** vérifier que `Board` est vide et `CurrentPlayer == 1`.

---

### `GetGame()`
**Rôle :** Retourne la partie en cours.  
**Signature :** `func GetGame() *Game`  
**Détails :** Accès centralisé à l’état du jeu stocké en global.  
**Complexité :** O(1)

---

### `(*Game) PlayColumn(col int)`
**Rôle :** Joue un pion dans une colonne et met à jour l’état de la partie.  
**Signature :** `func (g *Game) PlayColumn(col int) bool`  
**Détails :**  
- Vérifie si la colonne est valide.  
- Cherche la première case vide depuis le bas.  
- Place le pion du joueur actuel.  
- Appelle `checkWin` pour vérifier si le coup est gagnant.  
- Vérifie les matchs nuls (`isFull`).  
- Change de joueur automatiquement.  
**Complexité :** O(n) où *n = 6* (hauteur fixe de la grille).  
**Tests :**  
- coups valides / invalides  
- placement correct des pions  
- changement de joueur  

---

### `(*Game) checkWin(row, col int)`
**Rôle :** Vérifie si le dernier coup joué crée un alignement de 4 pions.  
**Signature :** `func (g *Game) checkWin(row, col int) bool`  
**Détails :** Analyse 4 axes :  
- horizontal  
- vertical  
- diagonale \  
- diagonale /  
**Complexité :** O(1) (grille de taille fixe, boucles bornées)  
**Tests :**  
- victoire horizontale, verticale, diagonale  
- non-victoires  

---

### `(*Game) isFull()`
**Rôle :** Indique si la grille est pleine.  
**Signature :** `func (g *Game) isFull() bool`  
**Détails :** Vérifie si la ligne supérieure contient encore des cases vides.  
**Complexité :** O(7) → considéré O(1).  
**Tests :** grille pleine / non pleine.


## Décisions d’architecture & compromis techniques

L’architecture du projet repose sur une structure simple et lisible afin de faciliter sa compréhension.  
La logique du jeu est centralisée dans la structure `Game`, ce qui permet d’assurer une gestion cohérente de l’état de la partie (grille, joueur courant, victoire, match nul). Le choix d’utiliser une variable globale pour stocker la partie en cours simplifie l’accès aux données depuis le serveur, même si cela limite la modularité ou la possibilité de gérer plusieurs parties simultanément.

## Qualité & tests

Aucun outil de vérification automatique n’a été utilisé (pas de linter type `go vet` ou `golint`).  
La validation du fonctionnement s’est faite manuellement, en testant directement le jeu via l’interface web.  
Aucun test unitaire automatisé n’a été écrit : les comportements (placement des pions, changement de joueur, détection de victoire et de match nul) ont été vérifiés au fur et à mesure de l’implémentation en exécutant le programme dans un navigateur.

## Limites connues & pistes d'amélioration

Les limites actuelles :

Voici quelques exemples pour t’inspirer (tu peux choisir ou ajouter les tiennes) :

une seule partie possible à la fois

pas de sauvegarde des parties

aucune IA / bot pour jouer

interface simple, pas responsive

pas d’animations

pas de test unitaire

architecture basique (pas d’API, pas de sessions)

code Go pas encore découpé en plusieurs fichiers






