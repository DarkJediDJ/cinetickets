package encryption

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// GetHash hashes password
func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
