package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"strings"

	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

//KeyStreams is a yaml configuration for a keystream used for encryption and decryption
type KeyStreams struct {
	Streams []struct {
		Name   string   `yaml:"name"`
		Stream []string `yaml:"stream"`
	} `yaml:"streams"`
}

var runeVals = map[rune]int{
	rune('A'): 0,
	rune('B'): 1,
	rune('C'): 2,
	rune('D'): 3,
	rune('E'): 4,
	rune('F'): 5,
	rune('G'): 6,
	rune('H'): 7,
	rune('I'): 8,
	rune('J'): 9,
	rune('K'): 10,
	rune('L'): 11,
	rune('M'): 12,
	rune('N'): 13,
	rune('O'): 14,
	rune('P'): 15,
	rune('Q'): 16,
	rune('R'): 17,
	rune('S'): 18,
	rune('T'): 19,
	rune('U'): 20,
	rune('V'): 21,
	rune('W'): 22,
	rune('X'): 23,
	rune('Y'): 24,
	rune('Z'): 25,
	rune(' '): 26,
	rune('.'): 27,
	rune(','): 28,
	rune(':'): 29,
	rune(';'): 30,
	rune('?'): 31,
}

var numVals = map[int]string{
	0:  "A",
	1:  "B",
	2:  "C",
	3:  "D",
	4:  "E",
	5:  "F",
	6:  "G",
	7:  "H",
	8:  "I",
	9:  "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "O",
	15: "P",
	16: "Q",
	17: "R",
	18: "S",
	19: "T",
	20: "U",
	21: "V",
	22: "W",
	23: "X",
	24: "Y",
	25: "Z",
	26: " ",
	27: ".",
	28: ",",
	29: ":",
	30: ";",
	31: "?",
}

var cardVals = map[string]int{
	"AC":  0,
	"2C":  1,
	"3C":  2,
	"4C":  3,
	"5C":  4,
	"6C":  5,
	"7C":  6,
	"8C":  7,
	"9C":  8,
	"10C": 9,
	"JC":  10,
	"QC":  11,
	"KC":  12,
	"AD":  13,
	"2D":  14,
	"3D":  15,
	"4D":  16,
	"5D":  17,
	"6D":  18,
	"7D":  19,
	"8D":  20,
	"9D":  21,
	"10D": 22,
	"JD":  23,
	"QD":  24,
	"KD":  25,
	"XY":  26,
	"XZ":  27,
	"AH":  28,
	"2H":  29,
	"3H":  30,
	"4H":  31,
	"5H":  32,
	"6H":  33,
	"7H":  34,
	"8H":  35,
	"9H":  36,
	"10H": 37,
	"JH":  38,
	"QH":  39,
	"KH":  40,
	"AS":  41,
	"2S":  42,
	"3S":  43,
	"4S":  44,
	"5S":  45,
	"6S":  46,
	"7S":  47,
	"8S":  48,
	"9S":  49,
	"10S": 50,
	"JS":  51,
	"QS":  52,
	"KS":  53,
}

func readKey(filename string, keyname string, verbose bool) []string {

	if "Default" == filename {
		usr, err := user.Current()
		check(err)
		filename = usr.HomeDir + "/.pontifex"
	}

	if verbose {
		fmt.Println("Reading keyfile: " + filename)
	}

	dat, err := ioutil.ReadFile(string(filename))
	check(err)
	var keyStreams KeyStreams
	err = yaml.Unmarshal(dat, &keyStreams)
	check(err)

	for index := range keyStreams.Streams {
		if keyStreams.Streams[index].Name == keyname {
			if verbose {
				fmt.Printf("Using Stream: %s\n%v\n", keyname, keyStreams.Streams[index].Stream)
			}
			return keyStreams.Streams[index].Stream
		}
	}
	log.Fatal("No stream found with the name: ", keyname)
	panic("")
}

func getCypherText(plaintext string, deck []string, verbose bool) (string, []string) {

	cryptext := ""
	cryptDeck := deck
	for _, plaintextChar := range plaintext {
		keystream, cutDeck := recursiveCut(cryptDeck)
		keystreamVal := cardVals[keystream]
		plaintextCharVal := runeVals[plaintextChar]
		cypherTextVal := (keystreamVal + plaintextCharVal) % len(runeVals)
		cypherTextChar := numVals[cypherTextVal]
		cryptext = cryptext + cypherTextChar
		if verbose {
			fmt.Printf("plaintextChar: %q\n", plaintextChar)
			fmt.Printf("plaintextCharVal: %d\n", plaintextCharVal)
			fmt.Printf("keystreamVal: %d\n", keystreamVal)
			fmt.Printf("cypherTextChar: %s\n", cypherTextChar)
			fmt.Printf("cypherTextVal: %d\n\n", cypherTextVal)
		}
		cryptDeck = cutDeck
	}

	return cryptext, cryptDeck
}

func getPlainText(cyphertext string, deck []string, verbose bool) (string, []string) {
	plaintext := ""
	cryptDeck := deck
	for _, cyphertextChar := range cyphertext {
		keystream, cutDeck := recursiveCut(cryptDeck)
		keystreamVal := cardVals[keystream]
		cyphertextCharVal := runeVals[cyphertextChar]
		plaintextTextVal := (cyphertextCharVal - keystreamVal)
		if verbose {
			fmt.Printf("original plaintextTextVal: %q\n", plaintextTextVal)
		}
		//nega mod
		for plaintextTextVal < 0 {
			plaintextTextVal += len(runeVals)
		}
		if verbose {
			fmt.Printf("modified plaintextTextVal: %q\n", plaintextTextVal)
		}
		plainTextChar := numVals[plaintextTextVal]
		plaintext = plaintext + plainTextChar
		if verbose {
			fmt.Printf("cyphertextChar: %q\n", cyphertextChar)
			fmt.Printf("keystreamVal: %d\n", keystreamVal)
			fmt.Printf("cyphertextCharVal: %d\n", cyphertextCharVal)
			fmt.Printf("plaintextText: %s\n", plainTextChar)
			fmt.Printf("plaintextTextVal: %d\n", plaintextTextVal)
			fmt.Printf("plaintext: %s\n\n", plaintext)
		}
		cryptDeck = cutDeck
	}

	return plaintext, cryptDeck
}

func recursiveCut(deck []string) (string, []string) {
	shifted := shiftJokersDown(deck)
	cut1 := triCut(shifted)
	cut2 := countCut(cut1)

	//count n number of cards into the deck where n is the first card if it is a joker then repeat the keystream algorithm
	firstCard := cut2[0]
	firstCardVal := cardVals[firstCard]

	keystream := cut2[firstCardVal%len(cut2)]

	if cardVals[keystream] == 27 || cardVals[keystream] == 26 {
		return recursiveCut(cut2)
	}
	return keystream, cut2
}

//Shift the A joker down place in the deck and the B joker down 2. Remember there is no way for the Jokers to wind up on the top of the deck.
func shiftJokersDown(deck []string) []string {
	aJoker, bJoker := -1, -1
	for index, value := range deck {
		if strings.Compare(value, "XY") == 0 {
			aJoker = index
		}
		if strings.Compare(value, "XZ") == 0 {
			bJoker = index
		}
	}

	if aJoker == -1 || bJoker == -1 {
		log.Fatal("Jokers not found, double check your deck")
	}

	//Use ,modulare arithmatic to keep us in bounds
	//The card after Joker
	temp := deck[(aJoker+1)%len(deck)]
	//Shift the AJoker down
	deck[(aJoker+1)%len(deck)] = deck[aJoker]
	//put temp back in the deck
	deck[aJoker] = temp

	//Do it again for the BJoker but twice
	//The card after Joker
	temp = deck[(bJoker+1)%len(deck)]
	deck[(bJoker+1)%len(deck)] = deck[bJoker]
	deck[bJoker] = temp

	temp = deck[(bJoker+2)%len(deck)]
	deck[(bJoker+2)%len(deck)] = deck[(bJoker+1)%len(deck)]
	deck[(bJoker+1)%len(deck)] = temp

	return deck
}

//Perform a tri cut on the deck by sectioning the deck between the two jokers and swapping the top and bottom portions
func triCut(deck []string) []string {
	//TODO find a way to stop copying stuff everywhere
	//Find Top and Bottom Joker, This isn't always the A joker on top
	top, bottom := -1, -1
	for index, value := range deck {
		//If it's a joker
		if strings.HasPrefix(value, "X") {
			//if it's the first joker we found assign it otherwise it's the bottom
			if top == -1 {
				top = index
			} else {
				bottom = index
				break
			}
		}
	}

	topDeck := deck[:top]
	//The jokers will stay in the mid section
	midDeck := deck[top : bottom+1]
	botDeck := deck[bottom+1:]
	retDeck := append(botDeck, midDeck...)
	return append(retDeck, topDeck...)
}

//Perform a count cut on the deck by observing the value of the card at the bottom of the deck. Remove that number of cards from the top of the deck and insert them just above the last card in the deck.
func countCut(deck []string) []string {
	//Get the value of the last card in the deck
	count := cardVals[deck[len(deck)-1]]
	if count == len(deck) {
		//if you are going to take the number of cards and put them in front of themselves...
		return deck
	}
	// fmt.Printf("Count: %d\n", count)
	//Take that number of cards
	topDeck := deck[:count]
	//to the end of the deck -1 card
	midDeck := deck[count : len(deck)-1]
	lastCard := deck[len(deck)-1]

	retDeck := append(midDeck, topDeck...)
	return append(retDeck, lastCard)
}

//Just a quick printer of the deck I'm using
func printDeck(deck []string) []string {
	for _, card := range deck {
		var val = 0
		if v, exists := cardVals[card]; exists {
			val = val + v
		} else {
			fmt.Printf("Key %s does not exist in map.\n", card)
			panic("Bad Card Value")
		}
		fmt.Printf("Card: %s, %d\n", card, val)
	}
	return deck
}
