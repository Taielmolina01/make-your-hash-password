package main 

import (
	"fmt"
	"os"
	"makeYourHashPassword/hashes"
	"makeYourHashPassword/commands"
	"makeYourHashPassword/errors"
	"makeYourHashPassword/createPassword"
    "strconv"
)

const (
	minLenValidArgs = 3
	lenValidArgsWithFlag = 1
	argIndexAlgorithm = 0
	argIndexPassword = 1
	argIndexMaxLengthPassword = 2
	argIndexRotatePassword = 3
	messageReturnPassword = "Your hashed password is: %s\n"
	messageHashNotAvailable = "Hash not available"
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

	if len(args) < minLenValidArgs {
		fmt.Println(errors.ErrorInvokingProgram{}.Error())
		return
	}

	algorithm := args[argIndexAlgorithm]
	password := args[argIndexPassword]
	lengthPassword, err := strconv.Atoi(args[argIndexMaxLengthPassword])
	if err != nil {
		fmt.Println("You must pass an integer to the length password's arg")
		return
	}
	
	isHash, h := hashes.IsHash(hashes.GetHashes(), algorithm)
	if isHash {
		rotations := 0
		if len(args) == minLenValidArgs+1 {
			var err error
			rotations, err = strconv.Atoi(args[argIndexRotatePassword])
			if err != nil {
				fmt.Println("You must pass an integer to the rotations arg")
				return
			}
		}
		hashed := createPassword.GetPassword(h, password, lengthPassword, rotations)
		fmt.Printf(messageReturnPassword, hashed)
	} else {
		fmt.Println(messageHashNotAvailable)
	}
}




