package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword will hash the password using the bcrypt algorithm
func HashPassword(rawPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 12)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes), nil
}

// CheckPasswordHash compares the password against the hashed password stored in memory
func CheckPasswordHash(rawPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err == nil
}

func main() {
	// Ask ask the user for their password and store it in the password variable
	var rawPassword string
	fmt.Println("Enter the password you want to encrypt")
	fmt.Scanln(&rawPassword)

	// run the password through the HashPassword Function and store it in the has variable
	hashedPassword, err := HashPassword(rawPassword)
	if err != nil {
		fmt.Println(err)
	}

	// show the user their raw password and the hashed version of their password
	fmt.Println("\nPassword:", rawPassword)
	fmt.Println("Hash:    ", hashedPassword)

	// ask the user to resubmit their password to check it against the hash
	fmt.Println("Please check your password against the hash")
	fmt.Scanln(&rawPassword)
	testPassword := CheckPasswordHash(rawPassword, hashedPassword)

	// display the confirmation result of the password
	fmt.Println("\nMatch:   ", testPassword)

}
