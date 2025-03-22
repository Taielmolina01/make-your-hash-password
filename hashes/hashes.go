package hashes

import (
	"crypto/sha256"
	"crypto/sha512"
	"crypto/hmac"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/blake2b"
	"hash"
)

const (
	SHA_256          = "SHA-256"
	SHA3_256         = "SHA3-256"
	BLAKE2s_256      = "BLAKE2s-256"
	SHA_512          = "SHA-512"
	SHA3_512         = "SHA3-512"
	BLAKE2b_512      = "BLAKE2b-512"
	HMAC_SHA256      = "HMAC-SHA256"
	SECRET_KEY_HMAC_SHA256 = "superSecret"
)

type hashType struct {
	name     string
	hashData func(string) string
}

func (h hashType) GetHashName() string {
	return h.name
}

func (h hashType) HashData(data string) string {
	return h.hashData(data)
}

func hashData(hasher hash.Hash, data string) string {
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

var hashes = []hashType{
	{"SHA-256", func(data string) string { return hashData(sha256.New(), data) }},
	{"SHA3-256", func(data string) string { return hashData(sha3.New256(), data) }},
	{"BLAKE2s-256", func(data string) string { h, _ := blake2s.New256(nil); return hashData(h, data) }},
	{"SHA-512", func(data string) string { return hashData(sha512.New(), data) }},
	{"SHA3-512", func(data string) string { return hashData(sha3.New512(), data) }},
	{"BLAKE2b-512", func(data string) string { h, _ := blake2b.New512(nil); return hashData(h, data) }},
	{"HMAC-SHA256", func(data string) string { return hashData(hmac.New(sha256.New, []byte(SECRET_KEY_HMAC_SHA256)), data) }},
}

// GetHashes returns a list of available hash functions
func GetHashes() []hashType {
	return hashes
}

func IsHash(hashes []hashType, possibleHash string) (bool, *hashType) {
	for _, h := range hashes {
		if h.GetHashName() == possibleHash {
			return true, &h
		}
	}
	return false, nil
}

