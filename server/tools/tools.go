// Package tools gives generic tools
package tools

import (
	"crypto/md5"
	"encoding/hex"
)

// Hash_password hash a password with md5 algorithm.
func Hash_password(password string) string {
	password_hasheur := md5.New()                              // hash the password
	password_hasheur.Write([]byte(password))                   // hash the password
	hashed_password := password_hasheur.Sum(nil)               // hash the password
	hashed_password_str := hex.EncodeToString(hashed_password) // hash the password

	return hashed_password_str
}
