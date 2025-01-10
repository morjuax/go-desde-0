package conditions

import (
	"fmt"

	"github.com/morjuax/go-desde-0/utils"
)

func LearningSwitch() {
	switch os := utils.GetOS(); os {
	case "linux": 
		utils.Print("This is linux")
	case "darwin":
		utils.Print("This is darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}