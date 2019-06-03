package main

import (
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readKey(filename string) {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	log.Print(string(dat))
}
