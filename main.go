package main

import (
	"fmt"
	"strings"
	"time"
)

// sus or not list
var suslist = "\x62\x65\x6e, \x42\x65\x6e, \x42\x72\x61\x6e\x64\x6f\x6e, \x62\x72\x61\x6e\x64\x6f\x6e, \x74\x65\x63\x68\x68\x69\x67\x68\x65\x73\x74, \x54\x65\x63\x68\x48\x69\x67\x68\x65\x73\x74, \x63\x72\x79\x70\x74\x6f\x66\x79\x72\x65, \x43\x72\x79\x70\x74\x6f\x66\x79\x72\x65"

func main() {
	// Start credit/info function
	motd()

	// Scan user input.
	fmt.Println()
	fmt.Println("Whats your name? ")
	var person string
	fmt.Scanln(&person)

	// Launch the susdenitfier function with given name
	susdentifier(person)
}

func susdentifier(person string) {
	fmt.Println()
	if strings.Contains(suslist, person) {
		fmt.Println("get out of here", person, "you sussy baka.")
		time.Sleep(5 * time.Second)
	} else {
		fmt.Println("Congratulations", person+"! You are NOT sus. Tell your family and friends and share this tool!")
		fmt.Println()
		fmt.Println("This tool will close in 5 seconds.")
		time.Sleep(5 * time.Second)
	}
}

func motd() {
	fmt.Println()
	fmt.Println("suschecker is a new Go written alternative to kirathetechie(jennasilicon)'s bensussy.py created by cryptofyre.")
	fmt.Println("Don't take it seriously, it was made for an inside joke xD")
	fmt.Println()
	fmt.Println("v1.0.0")
}
