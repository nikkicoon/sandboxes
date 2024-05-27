package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/big"
)

// CalculateHash returns a hash value given a string (key) and
// an error value, it returns 0 if an error occurs.
func CalculateHash(key string) string {
	// sha1 hash
	hash := sha1.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)
	// hexadecimal conversion
	hexSHA1 := hex.EncodeToString(hashBytes)
	// integer base16 conversion
	res, success := new(big.Int).SetString(hexSHA1, 16)
	if !success {
		panic("failed parsing big int from hex")
	}
	return res.String()
}

func main() {
	fmt.Println(CalculateHash("hamster"))
}
