package main 

import (
	"fmt"
	"os"
	"makeYourHashPassword/hashes"
	"makeYourHashPassword/commands"
	"makeYourHashPassword/errors"
)

const (
	lenValidArgs = 2
	lenValidArgsWithFlag = 1
	argIndexAlgorithm = 0
	argIndexPassword = 1
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(errors.ErrorInvokingProgram{}.Error())
		return
	}
	
	isCommand, c := commands.IsCommand(args[0])
	if len(args) == lenValidArgsWithFlag && isCommand {
		c.Invoke()
		return
	}

	if len(args) != lenValidArgs {
		fmt.Println(errors.ErrorInvokingProgram{}.Error())
		return
	}

	algorithm := args[argIndexAlgorithm]
	password := args[argIndexPassword]

	isHash, h := hashes.IsHash(hashes.GetHashes(), algorithm)
	if isHash {
		hashed := h.HashData(password)
		fmt.Printf("Your hashed password is: %s\n", hashed)
	} else {
		fmt.Println("Hash not available")
	}
}