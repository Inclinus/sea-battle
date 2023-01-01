package menu

import (
	"fmt"
	"sea-battle/internal/ip"

	"sea-battle/internal/board"
	"sea-battle/internal/boats"
	"sea-battle/internal/shots"
)

var aliases = make(map[string]ip.IP)

func DisplayChoices() {
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

func ManageAliases() {
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
			ip.DisplayAliases(&aliases)
		case 2:
			//display the ip of the player

		case 3:
			//Add an alias
			ip.AddAlias(&aliases, "192.168.12.2:4567", "Noam")
			ip.AddAlias(&aliases, "127.0.0.1:4567", "Thibaut")
		case 4:
			//remove an alias
			ip.RemoveAlias(&aliases, "Noam")
		case 5:
			fmt.Println("Retour au Menu Principal!")
			fmt.Println("------------------------------")

		default:
			fmt.Println("Votre choix doit etre entre 1 et 5 !")
		}
	}
}

func DisplayCredits() {
	fmt.Println("Nous souhaitons adresser nos remerciements les plus sincères à Monsieur Karraz qui nous a apporté son aide pour mener à  bien ce projet: " +
		"\"Jeu de Bataille Navale\"\n" +
		"Les Contributeurs au cours de ce projet : " +
		"- Anto BENEDETTI @opixelum\n" +
		"- Thibaut LULINSKI @KyatoNS\n" +
		"- Noam DE MASURE @Inclinus\n" +
		"- Charbel SALHAB @csalhabb\n" +
		"Merci !\n")
}

func DisplayRules() {
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

func ChooseOpponent() {
	fmt.Println("------------------------------")
	fmt.Println("Liste des joueurs :")
	ip.DisplayAliases(&aliases)
	fmt.Println("------------------------------")
	fmt.Println("Veillez entrer l'alias de l'adversaire : ")
	var selectedAlias string
	fmt.Scanf("%s\n", &selectedAlias)
	OpponentActions(selectedAlias)
}

func OpponentActions(selectedAlias string) {
	// PRINT DEBUG
	//ip.DisplayAlias(&aliases, selectedAlias)
	var ch int
	for ch != 4 {
		fmt.Println("Sous-Menu de choix d'action sur l'adversaire :\n" +
			"1 - Afficher son board\n" +
			"2 - Afficher son nombre de bateau\n" +
			"3 - Attaquer l'adversaire\n" +
			"4 - Quitter le Sous-Menu et retourner au Menu Principal\n" +
			"Quel est votre choix ?\n")

		fmt.Scanf("%d\n", &ch)

		switch ch {
		case 1:
			//display the board of the opponent
		case 2:
			//display the number of boats of the opponent

		case 3:
			//Attack the opponent
			fmt.Println("Veillez entrer la case à attaquer : ")
			var selectedCase string
			fmt.Scanf("%s\n", &selectedCase)
			pos := board.GetPositionFromString(selectedCase)
			//i, p := ip.GetIpOf("Noam", aliases)
			oppenentIp := ip.GetIpOf2(selectedAlias, &aliases)
			resultHit := shots.RequestHit(oppenentIp, pos)
			if resultHit == false {
				ChooseOpponent()
			}
		case 4:
			fmt.Println("Retour au Menu Principal!")
			fmt.Println("------------------------------")

		default:
			fmt.Println("Votre choix doit etre entre 1 et 5 !")
		}
	}
}

func DisplayMenu() {
	var choice int

	for choice != 7 {
		DisplayChoices()
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			//check the board state
			boats := boats.GenerateRandomBoats()

			// Get All Shots
			allShots := shots.GetShots()

			// Print board
			board.PrintBoard(boats, allShots)

		case 2:
			//Attack or start the game
			ChooseOpponent()
		case 3:
			ManageAliases()

		case 4:
			DisplayRules()

		case 5:
			//Statistics

		case 6:
			DisplayCredits()

		case 7:
			fmt.Println("Vous avez quitté le programme ! ")

		default:
			fmt.Println("Votre choix doit etre entre 1 et 7 !")
		}
	}
}
