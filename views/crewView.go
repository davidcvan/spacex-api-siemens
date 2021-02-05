package views

import (
	"fmt"
	"spacex-api-siemens/controllers"
)

func crewMenu() {
	var i int
	fmt.Printf("1 - Total crew members\n2 - Crew members by name\n3 - Total launches to space\n")
	fmt.Print("What info do you want: ")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Printf("There is a total of %v crew members\n", controllers.CrewCount())
		backOrExit()

	case 2:
		names := controllers.CrewNames()
		for _, name := range names {
			fmt.Println(name)
		}
		backOrExit()
	case 3:
		fmt.Printf("There are %v launches amongst %v crew\n", controllers.LaunchCount(), controllers.CrewCount())
		backOrExit()
	}
}
