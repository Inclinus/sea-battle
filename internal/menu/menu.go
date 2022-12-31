package menu

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sea-battle/internal/ip"

	"sea-battle/internal/board"
	"sea-battle/internal/boats"
	"sea-battle/internal/ip"
	"sea-battle/internal/shots"
)

// display menu choices
func DisplayChoices() {
	fmt.Println("Menu :\n" +
		"1 -  Voir l'état du board \n" +
		"2 -  Attaquer \n" +
		"3 -  Gérer les alias \n" +
		"4 -  Règles du jeu \n" +
		"5 -  Statistiques \n" +
		"6 -  Crédits \n" +
		"7 -  Quitter la session\n\n" +
		"Quel est votre choix ?")
}

func ManageAliases() {

	var ch int
	for ch != 5 {
		fmt.Println("Sous-Menu pour la Gestion des Alias :\n" +
			"1 - Afficher les Alias\n" +
			"2 - Afficher l’ip du joueur\n" +
			"3 - Ajouter un Alias\n" +
			"4 - Retirer un Alias\n" +
			"5 - Quitter le Sous-Menu et retourner au Menu Principal\n" +
			"Quel est votre choix ?\n")

		fmt.Scanf("%d\n", &ch)

		switch ch {
		case 1:
			ClearScreen()
			//display all the aliases
			fmt.Println("Voici la liste des alias : " + ip.displayAliases())

		case 2:
			ClearScreen()
			//display the ip of the player

		case 3:
			ClearScreen()
			//Add an alias

		case 4:
			ClearScreen()
			//remove an alias

		case 5:
			ClearScreen()
			fmt.Println("\nRetour au Menu Principal!\n")
			ip.SaveAlias()
		default:
			ClearScreen()
			fmt.Println("\nVotre choix doit etre entre 1 et 5 !\n")
		}
	}
}

// function that displays the credits of the game, or project
func DisplayCredits() {
	fmt.Println("\nNous souhaitons adresser nos remerciements les plus sincères à Monsieur Karraz qui nous " +
		"a apporté son aide pour mener à  bien ce projet: " +
		"\"Jeu de Bataille Navale\"\n" +
		"Les Contributeurs au cours de ce projet :\n" +
		"- Anto BENEDETTI @opixelum\n" +
		"- Thibaut LULINSKI @KyatoNS\n" +
		"- Noam DE MASURE @Inclinus\n" +
		"- Charbel SALHAB @csalhabb\n" +
		"Merci !\n")
}

// function that displays the rules of the games
func DisplayRules() {
	fmt.Println("\nLa bataille navale est un des jeux de société qui fait amuser petits et grands. " +
		"La bataille navale est idéale pour passer un moment en famille. " +
		"\nVoici les règles du jeu : \n\n" +
		"- La bataille navale se joue sur une grille faisant au moins 10x10.\n" +
		"- Le joueur doit deviner où se situent les bateaux adverses afin de les couler.\n" +
		"- Les bateaux devront être placés aléatoirement sur les grilles des joueurs, une case ne peut être " +
		"occupée que par un morceau de bateau à la fois.\n" +
		"- Il doit y avoir plusieurs bateaux présents sur le plateau.\n" +
		"- Par contre, nous ne jouons pas l’un après l’autre ici, mais en temps réel.\n" +
		"- N’importe quel joueur peut jouer plusieurs fois d’affilée et en continu sans attendre les actions des autres.\n" +
		"- Ce n’est pas un jeu au tour par tour.\n")
}

// clearScreen is a map from the operating system name to functions to execute
// the terminal commands to clear the screen on said OS
var clearScreen map[string]func()

// init initializes the clearScreen variable for MacOS, Linux, & Windows
func init() {
	clearScreen = make(map[string]func())

	clearScreen["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clearScreen["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clearScreen["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Clears the terminal window of the user if the operating system is supported
func ClearScreen() {
	function, exists := clearScreen[runtime.GOOS]
	if exists {
		function()
	}
}

// IsSupportedOS checks to see if the operating system that the user is running
// is able to have the terminal cleared
func IsSupportedOS() bool {
	_, exists := clearScreen[runtime.GOOS]
	return exists
}

// function that displays the menu
func DisplayMenu() {
	var choice int

	for choice != 7 {
		DisplayChoices()
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			ClearScreen()
			//check the board state
			boats := boats.GenerateRandomBoats()

			// Create an array of allShots
			var allShots []shots.Shot

			// Print board
			board.PrintBoard(boats, allShots)

		case 2:
			//Attack or start the game
			ClearScreen()
		case 3:
			ClearScreen()
			ManageAliases()

		case 4:
			ClearScreen()
			DisplayRules()

		case 5:
			ClearScreen()
			//Statistics

		case 6:
			ClearScreen()
			DisplayCredits()

		case 7:
			ClearScreen()
			fmt.Println("\nVous avez quitté le programme !\n")

		default:
			ClearScreen()
			fmt.Println("\nVotre choix doit etre entre 1 et 7 !\n")
		}
	}
}

func main() {
	DisplayMenu()
}
