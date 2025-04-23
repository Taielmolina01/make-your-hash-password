package createPassword

import (
	"makeYourHashPassword/hashes"
)

func GetPassword(h *hashes.HashType, password string, lengthPassword, rotations int) string {
	hashed := h.HashData(password)
	hashed = hashed[len(hashed)-rotations:] + hashed[:len(hashed)-rotations]
	return hashed[:lengthPassword]
}

func module(x, y int) int {
	return x % y
}