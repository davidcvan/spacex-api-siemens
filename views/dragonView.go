package views

import (
	"fmt"
	"spacex-api-siemens/controllers"
)

func dragonMenu() {
	var i int
	fmt.Printf("1 - Total number of Dragon capsules\n2 - Total payload sent to ISS\n3 - Days since first launch\n")
	fmt.Print("What info do you want: ")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Printf("There are currently %v dragon capsules\n", controllers.DragonCount())
		backOrExit()
	case 2:
		fmt.Printf("Dragon capsules took total of %v kg of payload to ISS\n", controllers.TotalPayload())
		backOrExit()
	case 3:
		m := controllers.DaysSinceLaunch()
		for name, days := range m {
			fmt.Printf("%v was first launched %v days ago\n", name, days)
		}
		backOrExit()
	}
}
