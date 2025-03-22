package commands

import (
	"fmt"
	"makeYourHashPassword/hashes"
)

const (
	helpCommand             = "-help"
	listAlgorithmsCommand   = "-l"
)

type command struct {
	name   string
	Invoke func()
}

func (c command) getCommandName() string {
	return c.name
}

var commands = []command{
	{
		name: helpCommand, 
		Invoke: func() {
			fmt.Println("You must call the executable like `executable algorithm password`")
			fmt.Println("Being `executable` the executable created when you compile main.go")
			fmt.Println("For algorithm you have these options: ")
			fmt.Println("And password is your password you want to hash")
		},
	},
	{
		name: listAlgorithmsCommand, 
		Invoke: func() {
			fmt.Println("The available hashing algorithms are: ")
			for _, v := range hashes.GetHashes() {
				fmt.Println(v.GetHashName())
			}
		},
	},
}

func GetCommands() []command {
	return commands
}

func IsCommand(commands []command, possibleCommand string) (bool, *command) {
	for _, c := range commands {
		if c.getCommandName() == possibleCommand {
			return true, &c
		}
	}
	return false, nil
}
