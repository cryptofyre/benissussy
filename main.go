package main

import (
	"fmt"
	"strings"
	"time"
)

// sus or not list
var suslist = "ben, Ben, Brandon, brandon, techhighest, TechHighest"

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
