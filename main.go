package main

import (
	"fmt"
)

func displayChoices() {

	fmt.Println("Menu :\n" +
		"1 -  Voir l'état du board \n" +
		"2 -  Attaquer \n" +
		"3 -  Gérer les alias \n" +
		"4 -  Règles du jeu \n" +
		"5 -  Statistiques \n" +
		"6 -  Crédits \n" +
		"7 -  Quitter la session\n\n" +
		"Quel est votre choix ? ")

}

func manageAliases() {
	fmt.Println("------------------------------")
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
			//display all the aliases
		case 2:
			//display the ip of the player
		case 3:
			//Add an alias
		case 4:
			//remove an alias
		case 5:
			fmt.Println("Retour au Menu Principal!")
			fmt.Println("------------------------------")
		default:
			fmt.Println("Votre choix doit etre entre 1 et 5 !")

		}
	}
}

func displayCredits() {

	fmt.Println("Nous souhaitons adresser nos remerciements les plus sincères à Monsieur Karraz qui nous a apporté son aide pour mener à  bien ce projet: " +
		"\"Jeu de Bataille Navale\"\n" +
		"Les Contributeurs au cours de ce projet : " +
		"- Anto BENEDETTI @opixelum\n" +
		"- Thibaut LULINSKI @KyatoNS\n" +
		"- Noam DE MASURE @Inclinus\n" +
		"- Charbel SALHAB @csalhabb\n" +
		"Merci !\n")

}

func displayRules() {

	fmt.Println("La bataille navale est un des jeux de société qui fait amuser petits et grands. La bataille navale est idéale pour passer un moment en famille. " +
		"\nVoici les règles du jeu : \n\n" +
		"- La bataille navale se joue sur une grille faisant au moins 10x10.\n" +
		"- Le joueur doit deviner où se situent les bateaux adverses afin de les couler.\n" +
		"- Les bateaux devront être placés aléatoirement sur les grilles des joueurs, une case ne peut être occupée que par un morceau de bateau à la fois.\n" +
		"- Il doit y avoir plusieurs bateaux présents sur le plateau.\n" +
		"- Par contre, nous ne jouons pas l’un après l’autre ici, mais en temps réel.\n" +
		"- N’importe quel joueur peut jouer plusieurs fois d’affilée et en continu sans attendre les actions des autres.\n" +
		"- Ce n’est pas un jeu au tour par tour.\n")

}

func displayMenu() {

	var choice int

	for choice != 7 {

		displayChoices()
		fmt.Scanf("%d\n", &choice)

		switch choice {

		case 1:
			//check the board state
		case 2:
			//Attack or start the game
		case 3:
			manageAliases()
		case 4:
			displayRules()
		case 5:
			//Statistics
		case 6:
			displayCredits()
		case 7:
			fmt.Println("Vous avez quitté le programme ! ")
		default:
			fmt.Println("Votre choix doit etre entre 1 et 7 !")
		}
	}

}

func main() {

	displayMenu()

}
