package createPassword

import (
	"makeYourHashPassword/hashes"
	"makeYourHashPassword/errors"
	"strconv"
	"fmt"
)

const (
	messageReturnPassword = "Your hashed password is: %s"
	messageHashNotAvailable = "Hash not available"
	MinLenValidArgs = 3
	argIndexAlgorithm = 0
	argIndexPassword = 1
	argIndexMaxLengthPassword = 2
	argIndexRotatePassword = 3
	errorMessageLengthPasswordNot
)

func getPassword(h *hashes.HashType, password string, lengthPassword, rotations int) string {
	hashed := h.HashData(password)
	hashed = hashed[len(hashed)-rotations:] + hashed[:len(hashed)-rotations]
	return hashed[:lengthPassword]
}

func module(x, y int) int {
	return x % y
}

func ManagePasswordRequest(args []string) string {
	algorithm := args[argIndexAlgorithm]
	password := args[argIndexPassword]
	lengthPassword, err := strconv.Atoi(args[argIndexMaxLengthPassword])
	if err != nil {
		return errors.ErrorTypeOfLengthPassword{}.Error()
	}
	
	isHash, h := hashes.IsHash(hashes.GetHashes(), algorithm)
	if !isHash {
		return messageHashNotAvailable
	} else {
		rotations := 0
		if len(args) == MinLenValidArgs + 1 {
			var err error
			rotations, err = strconv.Atoi(args[argIndexRotatePassword])
			if err != nil {
				return errors.ErrorTypeOfRotations{}.Error()
			}
		}
		hashed := getPassword(h, password, lengthPassword, rotations)
		return fmt.Sprintf(messageReturnPassword, hashed)
	}
}
