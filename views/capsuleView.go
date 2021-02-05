package views

import (
	"fmt"
	"spacex-api-siemens/controllers"
)

func capsuleMenu() {
	var i int
	fmt.Printf("There are curently %v capsules\n1 - Total land landings\n2 - Total water landings\n", controllers.CapsuleCount())
	fmt.Print("what info do you want: ")
	fmt.Scan(&i)
	switch i {
	case 1:
		fmt.Printf("There is a total of %v land landings\n", controllers.CapsuleLandLandings())
		backOrExit()

	case 2:
		fmt.Printf("There is a total of %v water landings\n", controllers.CapsuleWaterLandings())
		backOrExit()
	}

}
