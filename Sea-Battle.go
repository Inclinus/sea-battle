package main

import (
	"fmt"
	"os"
)

func DisplayMenu() {
	var choice int

	for true {
		fmt.Println("Menu :\n")
		fmt.Println("1-  Voir l'état du board \n")
		fmt.Println("2-  Attaquer \n")
		fmt.Println("3- Visualiser les alias \n")
		fmt.Println("4-  Règles du jeu \n")
		fmt.Println("5-  Statistiques \n")
		fmt.Println("6-  Crédits \n")
		fmt.Println("7-  Quitter la session \n")

		fmt.Println("Quel est votre choix ? \n")
		fmt.Scanf("%d\n", &choice)

		switch choice {

		case 1:
			//check the board state
		case 2:
			//Attack or start the game
		case 3:
			//Visualization of the alias
		case 4:
			fmt.Println("La bataille navale est un des jeux de société qui fait amuser petits et grands. La bataille navale est idéale pour passer un moment en famille. " +
				"\nVoici les règles du jeu : \n")
			fmt.Println("Pour jouer à la bataille navale, il vous faut un plateau de jeu où chacun dispose d’une grille numérotée de 1 à 10 horizontalement et annotée de A à J verticalement," +
				"ainsi que d'une flotte composée de quelques bateaux d'une à cinq cases de long. " +
				"\nIl ne faut pas que les participants voient la grille des adversaires.\n")
			fmt.Println("Au début du jeu, chaque joueur place ses bateaux sur sa grille, " +
				"en s'assurant que deux bateaux ne sont pas adjacents. L'autre représente la zone adverse, " +
				"où il cherchera à couler les bateaux de son adversaire.\n")
			fmt.Println("Chaque joueur, à son tour, annonce une case, et son adversaire lui répond si le tir tombe à l'eau ou au contraire s'il touche un bateau.\nDans ce dernier cas, " +
				"il annonce « touché » s'il reste des cases intactes au bateau ciblé, et « touché-coulé » si non.\n")
		case 5:
			//Statistics
		case 6:
			fmt.Println("Nous souhaitons adresser nos remerciements les plus sincères à Monsieur Karraz qui nous a apporté son aide pour mener à  bien ce projet.\n")
			fmt.Println("\"Jeu de Bataille Navale\"\n")
			fmt.Println("Les Contributeurs au cours de ce projet : \n")
			fmt.Println("- Anto BENEDETTI\n")
			fmt.Println("- Thibaut LULINSKI\n")
			fmt.Println("- Noam DE MASURE\n")
			fmt.Println("- Charbel SALHAB\n")
			fmt.Println("Merci !\n")
		case 7:
			os.Exit(0)
		default:
			fmt.Println("Votre choix doit etre entre 1 et 7 !")
		}
	}
}
func main() {

}
