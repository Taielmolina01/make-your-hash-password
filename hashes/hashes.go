package hashes

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
	"hash"
)

const (
	SHA_256                = "SHA-256"
	SHA3_256               = "SHA3-256"
	BLAKE2s_256            = "BLAKE2s-256"
	SHA_512                = "SHA-512"
	SHA3_512               = "SHA3-512"
	BLAKE2b_512            = "BLAKE2b-512"
	HMAC_SHA256            = "HMAC-SHA256"
	SHA1                   = "SHA-1"
	MD5                    = "MD5"
	SHA3_224               = "SHA3-224"
	BLAKE2s_128            = "BLAKE2s-128"
	SECRET_KEY_HMAC_SHA256 = "superSecret"
)

type HashType struct {
	name     string
	hashData func(string) string
}

func (h HashType) GetHashName() string {
	return h.name
}

func (h HashType) HashData(data string) string {
	return h.hashData(data)
}

func hashData(hasher hash.Hash, data string) string {
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

var hashes = []HashType{
	{SHA_256, func(data string) string { return hashData(sha256.New(), data) }},
	{SHA3_256, func(data string) string { return hashData(sha3.New256(), data) }},
	{BLAKE2s_256, func(data string) string { h, _ := blake2s.New256(nil); return hashData(h, data) }},
	{SHA_512, func(data string) string { return hashData(sha512.New(), data) }},
	{SHA3_256, func(data string) string { return hashData(sha3.New512(), data) }},
	{BLAKE2b_512, func(data string) string { h, _ := blake2b.New512(nil); return hashData(h, data) }},
	{HMAC_SHA256, func(data string) string { return hashData(hmac.New(sha256.New, []byte(SECRET_KEY_HMAC_SHA256)), data) }},
	{SHA1, func(data string) string { return hashData(sha1.New(), data) }},
	{MD5, func(data string) string { return hashData(md5.New(), data) }},
	{SHA3_224, func(data string) string { return hashData(sha3.New224(), data) }},
	{BLAKE2s_128, func(data string) string { h, _ := blake2s.New128(nil); return hashData(h, data) }},
}

func GetHashes() []HashType {
	return hashes
}

func IsHash(hashes []HashType, possibleHash string) (bool, *HashType) {
	for _, h := range hashes {
		if h.GetHashName() == possibleHash {
			return true, &h
		}
	}
	return false, nil
}
