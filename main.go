package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//var uinput []rune
var input []rune
var vertalingSecret []rune
var charnum [3]int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func importeerTekst() {
	content, err := ioutil.ReadFile("vertaal.txt")
	check(err)
	for i := range content {
		input = append(input, rune(content[i]))
	}
	fmt.Println(input)
}

func schrijfNaarOutput() {
	vertalingSecret = append(vertalingSecret, '🐋')
	fmt.Println(Alfabet)
	fmt.Println(vertalingSecret)
	file, err := os.Create("./vertaling.txt")
	check(err)
	file.WriteString(string(vertalingSecret))
	file.Close()
	file.Sync()
}

func indexNaarSymbool(number int) {
	var lengthSymbols int = len(Alfabet)
	//Eerste teken van alfabet wordt gebruikt als niet gevonden
	//Tweede teken van alfabet wordt gebruikt als einde zin
	charnum[0] = (number / lengthSymbols / lengthSymbols) + 2
	charnum[1] = ((number - ((charnum[0] - 2) * lengthSymbols * lengthSymbols)) / lengthSymbols) + 2
	charnum[2] = number - (((charnum[0] - 2) * lengthSymbols * lengthSymbols) + ((charnum[1] - 2) * lengthSymbols)) + 2

	var woord int = ((charnum[0] - 1) * lengthSymbols * lengthSymbols) + ((charnum[1] - 1) * lengthSymbols) + (charnum[2] - 1)
	fmt.Println(woord)
}

func SymboolNaarIndex(secret string) int {
	var index int = 0
	var lengthSymbols int = len(Alfabet)
	var cIndex [2]int
	var secretArray = [2]rune{rune(secret[0]), rune(secret[1])}
	for i := 0; i < 2; i++ {
		for j := 1; j < lengthSymbols; j++ {
			if secretArray[i] == Alfabet[j] {
				println(secretArray[i], Alfabet[i])
				cIndex[i] = j
			}
		}
	}
	index += cIndex[0] * lengthSymbols
	index += cIndex[1] * lengthSymbols
	return index
}

func zoekPositieNederlandsWoord(woord string) int {
	var positie int = -1
	for huidigWoord := 0; huidigWoord < len(Woordenboek); huidigWoord++ {
		if woord == Woordenboek[huidigWoord] {
			positie = huidigWoord
			goto result
		}
	}
result:
	return positie
}

func huidigWoord(startpositie int) string {
	var huidigWoord string
	for i := startpositie; input[i] != ' '; i++ {
		huidigWoord += string(input[i])
	}
	return huidigWoord
}

func vertaalNaarSecret() {
	//kijk naar volgend letter geen . of spatie
	//letter per letter toevoegen aan string en kijken als die overeenkomt met een woord
	var huidigWoord string
	var gevondenWoordPositie int
	for i := 0; i < len(input); i++ {
		if input[i] != ' ' && input[i] != '.' {
			huidigWoord += string(input[i])
			gevondenWoordPositie = zoekPositieNederlandsWoord(huidigWoord)
			if gevondenWoordPositie != -1 {
				indexNaarSymbool(gevondenWoordPositie)
				vertalingSecret = append(vertalingSecret, Alfabet[charnum[0]])
				vertalingSecret = append(vertalingSecret, Alfabet[charnum[1]])
				vertalingSecret = append(vertalingSecret, Alfabet[charnum[2]])
				huidigWoord = ""
			}
		} else if input[i] == ' ' {
			if gevondenWoordPositie == -1 {
				vertalingSecret = append(vertalingSecret, Alfabet[0])
				huidigWoord = ""
			}
			vertalingSecret = append(vertalingSecret, ' ')
		} else if input[i] == '.' {
			vertalingSecret = append(vertalingSecret, '.')
		}
	}
	schrijfNaarOutput()
}

func vertaalNaarNederlands() {
	schrijfNaarOutput()
}

func main() {
	var choice int
	for {
		fmt.Println("\n\n1. Laad woordenboek en alfabet in")
		fmt.Println("2. Laad vertaal.txt bestand in")
		fmt.Println("3. Vertaal ingeladen bestand naar Secret")
		fmt.Println("4. Vertaal ingeladen bestand naar Nederlands")
		fmt.Print("\nChoose nr: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			maakWoordenboek()
			importeerAlfabet()
		}
		if choice == 2 {
			importeerTekst()
		}
		if choice == 3 {
			vertaalNaarSecret()
		}
		if choice == 4 {
			vertaalNaarNederlands()
		}
	}
}
