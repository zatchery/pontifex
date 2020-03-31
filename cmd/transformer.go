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
	rune('A'): 1,
	rune('B'): 2,
	rune('C'): 3,
	rune('D'): 4,
	rune('E'): 5,
	rune('F'): 6,
	rune('G'): 7,
	rune('H'): 8,
	rune('I'): 9,
	rune('J'): 10,
	rune('K'): 11,
	rune('L'): 12,
	rune('M'): 13,
	rune('N'): 14,
	rune('O'): 15,
	rune('P'): 16,
	rune('Q'): 17,
	rune('R'): 18,
	rune('S'): 19,
	rune('T'): 20,
	rune('U'): 21,
	rune('V'): 22,
	rune('W'): 23,
	rune('X'): 24,
	rune('Y'): 25,
	rune('Z'): 26,
}

var numVals = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

var cardVals = map[string]int{
	"AC":  1,
	"2C":  2,
	"3C":  3,
	"4C":  4,
	"5C":  5,
	"6C":  6,
	"7C":  7,
	"8C":  8,
	"9C":  9,
	"10C": 10,
	"JC":  11,
	"QC":  12,
	"KC":  13,
	"AD":  14,
	"2D":  15,
	"3D":  16,
	"4D":  17,
	"5D":  18,
	"6D":  19,
	"7D":  20,
	"8D":  21,
	"9D":  22,
	"10D": 23,
	"JD":  24,
	"QD":  25,
	"KD":  26,
	"XY":  27,
	"XZ":  28,
	"AH":  29,
	"2H":  30,
	"3H":  31,
	"4H":  32,
	"5H":  33,
	"6H":  34,
	"7H":  35,
	"8H":  36,
	"9H":  37,
	"10H": 38,
	"JH":  39,
	"QH":  40,
	"KH":  41,
	"AS":  42,
	"2S":  43,
	"3S":  44,
	"4S":  45,
	"5S":  46,
	"6S":  47,
	"7S":  48,
	"8S":  49,
	"9S":  50,
	"10S": 51,
	"JS":  52,
	"QS":  53,
	"KS":  54,
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

func getCypherText(plaintext string, deck []string) (string, []string) {

	cryptext := ""
	cryptDeck := deck
	for _, plaintextChar := range plaintext {
		keystream, cutDeck := recursiveCut(cryptDeck)
		keystreamVal := cardVals[keystream]
		plaintextCharVal := runeVals[plaintextChar]
		cypherTextVal := (keystreamVal + plaintextCharVal) % len(runeVals)
		cypherTextChar := numVals[cypherTextVal]
		cryptext = cryptext + cypherTextChar
		cryptDeck = cutDeck
	}

	return cryptext, cryptDeck
}

func recursiveCut(deck []string) (string, []string) {
	shifted := shiftJokersDown(deck)
	cut1 := triCut(shifted)
	cut2 := countCut(cut1)

	//count n number of cards into the deck where n is the first card if it is a joker then repeat the keystream algorithm
	firstCard := cut2[0]
	firstCardVal := cardVals[firstCard]
	
	keystream := cut2[firstCardVal % len(cut2)]

	if cardVals[keystream] == 27 || cardVals[keystream] == 28 {
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
