package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

var Woordenboek []string
var Alfabet []rune

func importeerAlfabet() {
	content, err := ioutil.ReadFile("alfabet.txt")
	check(err)
	for i := range content {
		Alfabet = append(Alfabet, rune(content[i]))
	}
}

func maakWoordenboek() {
	file, err := os.Open("woordenboek.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		Woordenboek = append(Woordenboek, scanner.Text())
	}
}
