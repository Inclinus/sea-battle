package menu

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"sea-battle/internal/board"
	"sea-battle/internal/boats"
	"sea-battle/internal/ip"
	"sea-battle/internal/server"
	"sea-battle/internal/stats"
	"sea-battle/internal/utils"
	"strconv"
	"strings"
	"time"
)

var clearScreen map[string]func()

var ChallengeSentence string

// this function initializes the clearScreen variable for MacOS, linux and windows
func initClearScreenVariables() {
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

// this function will clear all the previously entered commands
// ONLY IF THE OS IS SUPPORTED
func ClearScreen() {
	function, exists := clearScreen[runtime.GOOS]
	if exists {
		function()
	}
}

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

// this function is used to add an alias, if the user wants to add many alias simultaneously, he must tap "o" as a yes
func AliasAddition() {

	var name string
	var userIp string
	var response string

	for true {
		fmt.Println("Quel est le nom de la personne que vous souhaitez ajouter ?")
		fmt.Scanf("%s\n", &name)

		fmt.Println("Quel est son IP ? (sous la forme 127.0.0.1:2555)?")
		fmt.Scanf("%s\n", &userIp)

		ip.AddAlias(userIp, name)

		a := ip.GetAlias()
		for key, value := range *a {
			if key == name {
				fmt.Printf("La personne %s a bien été ajoutée avec l'IP %s et le Port %s.\n", key, value.Ip, strconv.Itoa(int(value.Port)))
			}
		}

	myloop:
		for true {
			fmt.Println("Voulez-vous ajouter une autre personne ? (o/n)")
			fmt.Scanf("%s\n", &response)

			switch response {
			case "o":
				ClearScreen()
				break myloop
			case "n":
				ClearScreen()
				return
			default:
				ClearScreen()
				fmt.Println("Votre choix doit strictement etre soit oui \"o\" soit non \"n\"")
			}
		}
	}
}

// this function is used to search an alias, if the user wants to search many alias simultaneously, he must tap "o" as a yes
func searchAlias() {

	var name string
	var response string
	a := ip.GetAlias()

	for true {
		var found bool = false
		fmt.Println("De quelle personne voulez-vous voir l'IP ?")
		fmt.Scanf("%s\n", &name)

		for key := range *a {
			//fmt.Println(key)
			if name == key {
				ip.DisplayAlias(name)
				found = true
				break
			}
		}
		if found == false {
			fmt.Printf("%s n'a pas été trouvé.\n", name)
		}
	myloop:
		for true {
			fmt.Println("Voulez-vous voir l'IP d'une autre personne ? (o/n)")
			fmt.Scanf("%s\n", &response)

			switch response {
			case "o":
				ClearScreen()
				break myloop
			case "n":
				ClearScreen()
				return
			default:
				ClearScreen()
				fmt.Println("Votre choix doit strictement etre soit oui \"o\" soit non \"n\"")
			}
		}
	}
}

// this function is used to delete an alias, if the user wants to delete many alias simultaneously, he must tap "o" as a yes
func deleteFromAlias() {
	var name string
	var response string
	a := ip.GetAlias()

	for true {
		var found bool = false

		fmt.Println("Quel est le nom de la personne que vous souhaitez supprimer ?")
		fmt.Scanf("%s\n", &name)

		for key := range *a {
			//fmt.Println(key)
			if name == key {
				ip.RemoveAlias(name)
				found = true
				break
			}
		}
		if found == false {
			fmt.Printf("%s n'a pas été trouvé, il n'a donc pas été supprimé.\n", name)
		}
	myloop:
		for true {
			fmt.Println("Voulez-vous supprimer une autre personne ? (o/n)")
			fmt.Scanf("%s\n", &response)

			switch response {
			case "o":
				ClearScreen()
				break myloop
			case "n":
				ClearScreen()
				return
			default:
				ClearScreen()
				fmt.Println("Votre choix doit strictement etre soit oui \"o\" soit non \"n\"")
			}
		}
	}
}

// the sub-menu for managing aliases
func ManageAliases() {

	var ch int

	for ch != 5 {
		fmt.Println("Sous-Menu pour la Gestion des Alias :\n" +
			"1- Afficher les Alias\n" +
			"2- Afficher l’ip d'un joueur\n" +
			"3- Ajouter un Alias\n" +
			"4- Retirer un Alias\n" +
			"5- Quitter le Sous-Menu et retourner au Menu Principal\n\n" +
			"Quel est votre choix ?")

		fmt.Scanf("%d\n", &ch)

		switch ch {
		case 1:
			ClearScreen()
			ip.DisplayAliases(true)
			fmt.Print("Appuyez sur Entrée pour revenir à la gestion des alias...")
			fmt.Scanln()
			ClearScreen()
		case 2:
			ClearScreen()
			ip.DisplayAliases(false)
			searchAlias()
		case 3:
			ClearScreen()
			AliasAddition()
		case 4:
			ClearScreen()
			ip.DisplayAliases(false)
			deleteFromAlias()
		case 5:
			ClearScreen()
			fmt.Println("\nRetour au Menu Principal!\n")
			ip.SaveAlias()
		default:
			ClearScreen()
			fmt.Println("\nVotre choix doit etre entre 1 et 5!\n")
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
	fmt.Print("Appuyez sur Entrée pour revenir au menu principal...")
	fmt.Scanln()
	ClearScreen()
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
	fmt.Print("Appuyez sur Entrée pour revenir au menu principal...")
	fmt.Scanln()
	ClearScreen()
}

func DisplayStats() {
	fmt.Printf("📈 Statistiques 📈\n\n")
	stat := stats.GetStats()
	fmt.Println("Parties jouées  :", stat.GamesWon+stat.GamesLost)
	fmt.Println("Parties gagnées :", stat.GamesWon)
	fmt.Println("Parties perdues :", stat.GamesLost)
	fmt.Println("Tir effectués   :", stat.ShotsHit+stat.ShotsMissed)
	fmt.Println("Tir réussis     :", stat.ShotsHit)
	fmt.Println("Tir ratés       :", stat.ShotsMissed)
	fmt.Println("\nAppuyez sur Entrée pour revenir au menu principal...")
	fmt.Scanln()
	ClearScreen()
}

func ChooseOpponent() {
	ip.DisplayAliases(true)
	fmt.Println("Veuillez entrer l'alias de l'adversaire ou 'exit' pour quitter : ")
	var selectedAlias string
	fmt.Scanf("%s\n", &selectedAlias)
	if selectedAlias == "exit" {
		ClearScreen()
		fmt.Println("\nRetour au Menu Principal!\n")
		return
	}
	ResultAliasIsExist := ip.AliasIsExist(selectedAlias)
	if ResultAliasIsExist {
		if ip.IsConnected(ip.GetIpOf(selectedAlias)) {
			OpponentActions(selectedAlias)
		} else {
			fmt.Println("L'adversaire n'est pas connecté !")
		}
	} else {
		fmt.Println("L'alias n'existe pas !")
		ChooseOpponent()
	}
}

func OpponentActions(selectedAlias string) {
	ClearScreen()
	var ch int
	for ch != 4 {
		fmt.Println("Sous-Menu de choix d'action sur " + selectedAlias + " :\n" +
			" 1 - Afficher son board\n" +
			" 2 - Afficher son nombre de bateau\n" +
			" 3 - Attaquer l'adversaire\n" +
			" 4 - Quitter le Sous-Menu et retourner au Menu Principal\n" +
			"Quel est votre choix ?\n")

		fmt.Scanf("%d\n", &ch)

		switch ch {
		case 1:
			ClearScreen()
			//display the board of the opponent
			board.RequestBoard(ip.GetIpOf(selectedAlias))
		case 2:
			ClearScreen()
			enemyIp := ip.GetIpOf(selectedAlias)

			port := strconv.Itoa(int(enemyIp.Port))
			url := "http://" + enemyIp.Ip + ":" + port + "/boats"

			client := http.Client{
				Timeout: 2 * time.Second,
			}

			resp, err := client.Get(url)
			if err != nil {
				fmt.Println("Une erreur est survenue.")
				return
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Une erreur est survenue.")
				return
			}
			result := string(body)
			fmt.Println("Il reste " + result + " bateau(x) à " + selectedAlias + ".\n")
		case 3:
			ClearScreen()
			//Attack the opponent
			board.RequestBoard(ip.GetIpOf(selectedAlias))
			isCaseValid := false
			fmt.Println("🎯 " + strings.ToUpper(selectedAlias) + " 🎯")
			var pos utils.Position
			for !isCaseValid {
				fmt.Println("Veuillez entrer la case à attaquer : (format : A1 ou a1) ")
				var selectedCase string
				fmt.Scanf("%s\n", &selectedCase)
				pos = board.GetPositionFromString(selectedCase)
				if pos.X == 0 || pos.Y == 0 || pos.Y > 10 {
					fmt.Println("La case entrée n'est pas valide !")
				} else {
					isCaseValid = true
				}
			}
			oppenentIp := ip.GetIpOf(selectedAlias)
			resultHit := board.RequestHit(oppenentIp, pos)
			if resultHit == false {
				ChooseOpponent()
				ch = 4
			}
			fmt.Print("Appuyez sur Entrée pour revenir aux actions sur " + selectedAlias + "...")
			fmt.Scanln()
			ClearScreen()
		case 4:
			ClearScreen()
			fmt.Println("Retour au Menu Principal!")
			fmt.Println("------------------------------")
		default:
			ClearScreen()
			fmt.Println("Votre choix doit etre entre 1 et 5 !")
		}
	}
}

func InitMenu() {
	initClearScreenVariables()
	var boatsBoard [5]boats.Boat

	choice := "CHOICE"
	for choice != "O" {
		ClearScreen()
		boatsBoard = boats.GenerateRandomBoats()
		board.PrintBoard(boatsBoard, false, ChallengeSentence)
		fmt.Println("Voici votre board, est-ce qu'il vous satisfait ? (O/N)")
		fmt.Scanf("%s\n", &choice)
	}
	ChallengeSentence = " "
	satisfied := "N"
	for satisfied != "O" {
		fmt.Println("Veuillez choisir une phrase pour défier vos adversaires !")
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		ChallengeSentence = input
		fmt.Println("Voici votre phrase :\n" + ChallengeSentence)
		fmt.Println("Voulez-vous utiliser cette phrase de défi ? (O/N)")
		fmt.Scanf("%s\n", &satisfied)
	}

	board.InitBoatsBoard(boatsBoard)
	go server.LaunchServer(ChallengeSentence)

	displayMenu()
}

func displayMenu() {
	var choice int
	ClearScreen()
	for choice != 7 {
		DisplayChoices()
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			// Print board
			ClearScreen()
			board.PrintBoard(board.GetBoatsBoard(), false, ChallengeSentence)
			//DEBUG
			//test := board.PrintBoard2(board.GetBoatsBoard(), false)
			//fmt.Println(test)
		case 2:
			//Attack or start the game
			ClearScreen()
			ChooseOpponent()
		case 3:
			ClearScreen()
			ManageAliases()
		case 4:
			ClearScreen()
			DisplayRules()
		case 5:
			ClearScreen()
			DisplayStats()
		case 6:
			ClearScreen()
			DisplayCredits()
		case 7:
			ClearScreen()
			fmt.Println("\nVous avez quitté le programme !\n")
		default:
			ClearScreen()
			fmt.Println("\nVotre choix doit etre entre 1 et 8 !\n")
		}
	}
}
