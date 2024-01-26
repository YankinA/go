package main

import (
	"fmt"

	greeting1 "github.com/YankinA/go/greeting"
	greeting2 "github.com/YankinA/go/greeting/v2"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(greeting1.Get())
	fmt.Println(greeting2.Get())
	logrus.Println("Hello, Hexlet!")
	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")
}
