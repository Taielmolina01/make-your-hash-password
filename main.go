package main 

import (
	"fmt"
	"os"
	"makeYourHashPassword/commands"
	"makeYourHashPassword/errors"
	"makeYourHashPassword/createPassword"
)

const (
	lenValidArgsWithFlag = 1
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(errors.ErrorInvokingProgram{}.Error())
		return
	}
	commandsAvailable := commands.InitializeCommands()
	isCommand, c := commands.IsCommand(args[0], commandsAvailable)
	if len(args) == lenValidArgsWithFlag && isCommand {
		c.Invoke()
		return
	}

	if len(args) < createPassword.MinLenValidArgs {
		fmt.Println(errors.ErrorInvokingProgram{}.Error())
		return
	}

	fmt.Println(createPassword.ManagePasswordRequest(args))
}