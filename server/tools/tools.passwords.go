// Package tools gives generic tools
package tools

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"unicode"
)

// Hash_password hash a password with md5 algorithm.
func Hash_password(password string) string {
	password_hasheur := md5.New()
	password_hasheur.Write([]byte(password))
	hashed_password := password_hasheur.Sum(nil)
	hashed_password_str := hex.EncodeToString(hashed_password)

	return hashed_password_str
}

// Password_is_valid_v1 check if a password is valid (v1).
func Password_is_valid_v1(password string) error {
	if len(password) < 8 || len(password) > 30 {
		return errors.New("password must be between 8 and 30 characters")
	}

	lower := false
	upper := false
	number := false
	special := false

	for _, character := range password {
		switch {
		case unicode.IsLower(character):
			lower = true
		case unicode.IsUpper(character):
			upper = true
		case unicode.IsNumber(character):
			number = true
		case unicode.IsPunct(character) || unicode.IsSymbol(character):
			special = true
		default:
			// If the password contains an invalid character.
			return errors.New("password should only contain letters, numbers and special characters")
		}
	}

	if !lower {
		// If the password does not contain at least one lowercaser.
		return errors.New("password must contain at least one lowercase letter")
	}
	if !upper {
		// If the password does not contain at least one uppercase.
		return errors.New("password must contain at least one uppercase letter")
	}
	if !number {
		// If the password does not contain at least one number.
		return errors.New("password must contain at least one number")
	}
	if !special {
		// If the password does not contain at least one special character.
		return errors.New("password must contain at least one special character")
	}

	return nil
}
