// Dimensions du plateau
const rows = 7;
const cols = 8;

// Joueur courant : 1 = rouge, 2 = jaune
let currentPlayer = 1;

// Plateau côté client
const boardDiv = document.getElementById('board');
const board = [];

// Création du plateau
function createBoard() {
    boardDiv.innerHTML = ""; // Vide le plateau
    for (let r = 0; r < rows; r++) {
        const rowDiv = document.createElement('div');
        rowDiv.classList.add('row');
        board[r] = [];
        for (let c = 0; c < cols; c++) {
            const cellDiv = document.createElement('div');
            cellDiv.classList.add('cell');
            cellDiv.dataset.row = r;
            cellDiv.dataset.col = c;
            // Clic sur une cellule → place le jeton dans la colonne
            cellDiv.addEventListener('click', () => playColumn(c));
            rowDiv.appendChild(cellDiv);
            board[r][c] = cellDiv;
        }
        boardDiv.appendChild(rowDiv);
    }
}

// Fonction pour jouer dans une colonne
function playColumn(col) {
    for (let r = rows - 1; r >= 0; r--) {
        const cell = board[r][col];
        if (!cell.classList.contains('player1') && !cell.classList.contains('player2')) {
            cell.classList.add(currentPlayer === 1 ? 'player1' : 'player2');
            currentPlayer = currentPlayer === 1 ? 2 : 1;
            updateCurrentPlayerDisplay();
            break;
        }
    }
}

// Mise à jour du joueur actuel
function updateCurrentPlayerDisplay() {
    const display = document.getElementById('currentPlayer');
    display.textContent = currentPlayer === 1 ? 'Joueur 1 (Rouge)' : 'Joueur 2 (Jaune)';
    display.className = currentPlayer === 1 ? 'player1' : 'player2';
}

// Réinitialisation du plateau
function resetBoard() {
    createBoard();
    currentPlayer = 1;
    updateCurrentPlayerDisplay();
}

// Initialisation au chargement
createBoard();
