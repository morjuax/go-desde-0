package keyboard

import (
	"bufio"
	"os"
	"strconv"

	"github.com/morjuax/go-desde-0/utils"
)

var number1 int
var number2 int
var leyenda string
var err error


func IngressNumbers() {
  scanner := bufio.NewScanner(os.Stdin)

  utils.Print("Input number 1: ")
  if scanner.Scan() {
    number1, err = strconv.Atoi(scanner.Text())
    if err != nil {
      panic("Data incorrect " + err.Error())
    }
  }

  utils.Print("Input number 2: ")
  if scanner.Scan() {
    number2, err = strconv.Atoi(scanner.Text())
    if err != nil {
      panic("Data incorrect " + err.Error())
    }
  }

  utils.Print("Input leyenda 2: ")
  if scanner.Scan() {
    leyenda = scanner.Text()
  }

  utils.Print(leyenda, number1 * number2)
}