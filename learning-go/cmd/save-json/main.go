/*
Save JSON
*/

package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	person := Person{Name: Name{Family: "Newmarch", Personal: "Jan"}, Email: []Email{Email{Kind: "home", Address: "jan@newmarch.name"}, Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	saveJSON("person.json", person)
}

func saveJSON(fileName string, key interface{}) {

	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {

	log.Fatal(err)
	os.Exit(1)
}
