package main 

import (
	"fmt"
	"os"
	"makeYourHashPassword/hashes"
	"makeYourHashPassword/commands"
	"makeYourHashPassword/errors"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(errors.ErrorInvokingProgram{}.Error())
		return
	}
	
	isCommand, c := commands.IsCommand(commands.GetCommands(), args[0])
	if len(args) == 1 && isCommand {
		c.Invoke()
		return
	}

	if len(args) != 2 {
		fmt.Println(errors.ErrorInvokingProgram{}.Error())
		return
	}

	algorithm := args[0]
	password := args[1]

	isHash, h := hashes.IsHash(hashes.GetHashes(), algorithm)
	if isHash {
		hashed := h.HashData(password)
		fmt.Printf("Your hashed password is: %s\n", hashed)
	} else {
		fmt.Println("Hash not available")
	}
}