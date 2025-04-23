package createPassword

import (
	"fmt"
	"makeYourHashPassword/errors"
	"makeYourHashPassword/hashes"
	"strconv"
)

const (
	messageReturnPassword     = "Your hashed password is: %s"
	MinLenValidArgs           = 3
	argIndexAlgorithm         = 0
	argIndexPassword          = 1
	argIndexMaxLengthPassword = 2
	argIndexRotatePassword    = 3
)

func getPassword(h *hashes.HashType, password string, lengthPassword, rotations int) string {
	hashed := h.HashData(password)
	hashed = hashed[len(hashed)-module(rotations, len(hashed)):] + hashed[:len(hashed)-module(rotations, len(hashed))]
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
		return errors.ErrorHashNotAvailable{}.Error()
	} else {
		rotations := 0
		if len(args) == MinLenValidArgs+1 {
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
