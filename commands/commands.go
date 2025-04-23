package commands

import (
	"fmt"
	"makeYourHashPassword/hashes"
)

const (
	helpCommand           = "-h"
	listAlgorithmsCommand = "-l"
	helpCommandMessage    = `You must call the executable like 'executable algorithm password maxLengthPassword optional[numberOfRotations]'
Being 'executable' the executable created when you compile main.go (by default is 'createYourPassword')

For algorithm you have these options:
%s

Password is your password you want to hash
MaxLengthPassword is the number of characters you want for your password
And the optional arg numberOfRotations is the number of rotations to the right that you want for your password. By default is 0`
	listAlgorithmsMessage = `The available hashing algorithms are:
%s`
)

type command struct {
	name    string
	message string
	Invoke  func()
}

func (c command) getCommandName() string {
	return c.name
}

func InitializeCommands() []command {
	return []command{
		initializeHelpCommand(),
		initializeListAlgorithmsCommand(),
	}
}

func initializeHelpCommand() command {
	availableAlgorithms := getHashesString()
	messageHelp := fmt.Sprintf(helpCommandMessage, availableAlgorithms)
	cmd := command{
		name:    helpCommand,
		message: messageHelp,
	}
	cmd.Invoke = func() {
		fmt.Println(cmd.message)
	}
	return cmd
}

func initializeListAlgorithmsCommand() command {
	availableAlgorithms := getHashesString()
	messageList := fmt.Sprintf(listAlgorithmsMessage, availableAlgorithms)
	cmd := command{
		name:    listAlgorithmsCommand,
		message: messageList,
	}
	cmd.Invoke = func() {
		fmt.Println(cmd.message)
	}
	return cmd
}

func getHashesString() string {
	var hashesString string
	for _, v := range hashes.GetHashes() {
		hashesString += v.GetHashName() + "\n"
	}
	return hashesString[:len(hashesString)-1]
}

func IsCommand(possibleCommand string, commands []command) (bool, *command) {
	for _, c := range commands {
		if c.getCommandName() == possibleCommand {
			return true, &c
		}
	}
	return false, nil
}
