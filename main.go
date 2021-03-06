package main

import (
	"fmt"
	"log"

	"crypto/hmac"
	"crypto/sha512"

	"golang.org/x/crypto/bcrypt"
)

var key []byte

func main() {
	for i := 1; 0 <= 64; i++ {
		key = append(key, byte(i))
	}

	pass := "123456789"

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	fmt.Println(pass)
	fmt.Println(hashedPass)

	log.Println("Logged in!")
}

// how to hash a password & store it in your database
func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error in hashPassword while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

// pass in password from user & get hashed password for user to compare/see if they match
func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Error in comparePassword while comparing password: %w", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)

	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while hashing message: %w", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)

	if err != nil {
		return false, fmt.Errorf("Error in checkSig while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSig, sig)

	return same, nil
}
