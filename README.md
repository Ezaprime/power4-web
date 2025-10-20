# Power 4 (Connect 4) – Version Go + Templates HTML

Power 4 est un jeu de Puissance 4 développé en Go (Golang) avec un affichage en HTML/CSS.  
Deux joueurs s’affrontent en déposant des pions, tour par tour, jusqu’à ce qu’un joueur aligne quatre jetons ou que le match se termine par une égalité.

---

## Fonctionnalités

- Entrée des prénoms des joueurs  
- Choix de la difficulté (taille du plateau variable)  
- Affichage dynamique du plateau  
- Indication du joueur en cours  
- Détection de victoire ou match nul  
- Page de résultat dédiée  
- Boutons Rejouer et Accueil  

---

## Technologies utilisées

| Technologie | Rôle |
|------------|------|
| Go (Golang) | Back-end et logique du jeu |
| HTML Templates | Pages d’affichage (home, play, result) |
| CSS | Mise en forme visuelle |
| JavaScript | Gestion des actions sans rechargement complet |

---

## Structure du projet

├── main.go
├── power4/
│ ├── server.go
│ ├── handler.go
│ ├── game.go
│
├── templates/
│ ├── home.html
│ ├── play.html
│ ├── result.html
│
└── static/
└── style.css
