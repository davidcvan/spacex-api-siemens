package views

import (
	"fmt"
	"os"
)

func MainMenu() {
	var i int
	fmt.Printf("Welcome to Spacex insights\nChoose what you want to get insights about:\n1 - Capsules\n2 - Crew\n3 - Dragon capsules\n")
	fmt.Print("What do you want more info about: ")
	fmt.Scan(&i)
	switch i {
	case 1:
		capsuleMenu()
	case 2:
		crewMenu()
	case 3:
		dragonMenu()
	}
}

func backOrExit() {
	var i int
	fmt.Printf("9 - Back to main menu\n0 - Exit\n")
	fmt.Print("what do you want to do: ")
	fmt.Scan(&i)
	switch i {
	case 9:
		MainMenu()
	case 0:
		os.Exit(3)
	}

}
