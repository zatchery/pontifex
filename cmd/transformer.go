package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

type KeyStreams struct {
	Streams []struct {
		Name   string   `yaml:"name"`
		Stream []string `yaml:"stream"`
	} `yaml:"streams"`
}

func readKey(filename string, keyname string, verbose bool) []string {

	// files, err := filepath.Glob("*")
	// check(err)
	// fmt.Println(files) // contains a list of all files in the current directory
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

func translate(text string, deck []string) string {

	var cardVals = map[string]int{
		"A":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"J":  11,
		"Q":  12,
		"K":  13,
		"D":  0,
		"C":  13,
		"H":  26,
		"S":  39,
		"XY": 53,
		"XZ": 54}
	for _, card := range deck {
		var val = 0
		for k, v := range cardVals {
			if strings.Contains(card, k) {
				val = val + v
			}
		}
		fmt.Printf("Card: %s, %d\n", card, val)
	}
	/**
	The deck will be assumed to be a circular array, meaning that should a card ever need to advance below the bottom card in the deck, it will simply rotate back to the top (in other words, the first card follows the last card).
	Arrange the deck of cards face-up according to a specific key. This is the most important part as anyone who knows the deck's starting value can easily generate the same values from it. How the deck is initialized is up to the recipients, shuffling the deck perfectly randomly is preferable, although there are many other methods. For this example, the deck will simply start at 1 and count up by 3's, modulo 28. Thus the starting deck will look like this:
	1 4 7 10 13 16 19 22 25 28 3 6 9 12 15 18 21 24 27 2 5 8 11 14 17 20 23 26
	Locate the first joker (value 27) and move it down the deck by one place, basically just exchanging with the card below it. Notice that if it is the last card, it becomes the second card. There is no way to become the first card. The deck now looks like this:
	1 4 7 10 13 16 19 22 25 28 3 6 9 12 15 18 21 24 2 27 5 8 11 14 17 20 23 26
	Locate the second joker (value 28) and move it down the deck by two places. Notice that if it is the second to last card, it becomes the second card by wrapping around. If it is the last card, it becomes the third card. There is no way to become the first card.
	1 4 7 10 13 16 19 22 25 3 6 28 9 12 15 18 21 24 2 27 5 8 11 14 17 20 23 26
	Perform a triple-cut on the deck. That is, split the deck into three sections. Everything above the top joker (which, after several repetitions, may not necessarily be the first joker) and everything below the bottom joker will be exchanged. The jokers themselves, and the cards between them, are left untouched.
	5 8 11 14 17 20 23 26 28 9 12 15 18 21 24 2 27 1 4 7 10 13 16 19 22 25 3 6
	Observe the value of the card at the bottom of the deck, if the card is either joker let the value just be 53. Take that number of cards from the top of the deck and insert them back to the bottom of the deck just above the last card.
	23 26 28 9 12 15 18 21 24 2 27 1 4 7 10 13 16 19 22 25 3 5 8 11 14 17 20 6
	Now, look at the value of the top card. Looking at the deck used above, the top card would be 23. Count this many places below that top card and take the value of that card as the next value in the keystream, in this example it would be the 24th card, which is 11. If the card counted to is either joker, do not add anything to the keystream. (Note that no cards are changing places in this step, this step simply determines the value).
	Repeat steps 2 through 6 for as many keystream values as required.
	It is worth noting that because steps 2 and 3 have the wrap around feature, there are two configurations that can lead to the same configuration on a step. For instance, when the little joker is either on the bottom of the deck or on the top of the deck it will become the second card after step 2. This means the algorithm is not always reversible.[2]
	**/

	return ""
}
