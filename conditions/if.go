package conditions

import (
	"github.com/morjuax/go-desde-0/utils"
)

func LearningIf() {
		// clasic
	// os := runtime.GOOS
	// if os=="darwin" {
	// 	utils.Print("Is darwin IOS")
	// } else {
	// 	utils.Print("Is not darwin IOS")
	// }

	// Performace
	if os := utils.GetOS(); os=="darwin" {
		utils.Print("Is darwin IOS")
	} else {
		utils.Print("Is not darwin IOS")
	}
}