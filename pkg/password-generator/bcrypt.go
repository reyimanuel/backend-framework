package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Package bcrypt provides a way to hash passwords using the bcrypt algorithm.
func Generate(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// This function mainly for generating user password hash for the first time. Not for register purposes.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	hash, err := Generate(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generated hash:", hash)
}
