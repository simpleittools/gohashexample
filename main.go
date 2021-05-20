package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// HashPassword will hash the password using the bcrypt algorithm
func HashPassword(rawPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 12)
	if err != nil {
		fmt.Println(err)
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares the password against the hashed password stored in memory
func CheckPasswordHash(rawPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err == nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return

	}
	fmt.Fprintf(w, "hello")
}

func main() {
	port := ":7080"
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
	// https://blog.logrocket.com/creating-a-web-server-with-golang/

	http.HandleFunc("/", indexHandler)

	fmt.Printf("Server is running at http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
