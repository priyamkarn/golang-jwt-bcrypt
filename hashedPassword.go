package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// DefaultCost means 2^10 = 1024 rounds of hashing are done

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(passwordSentByUser string, hashedPasswordInDB string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswordInDB), []byte(passwordSentByUser))

	return err == nil
	//IF THE PASSWORD MATCHES, IT RETURNS nil
	// RETURNS true IF PASSWORD MATCHES, ELSE false
}

func main() {
    // Test password
    password := "mySecurePassword123"
    
    // Hash the password
    hashedPassword, err := HashPassword(password)
    if err != nil {
        fmt.Printf("Error hashing password: %v\n", err)
        return
    }
    fmt.Printf("Original password: %s\n", password)
    fmt.Printf("Hashed password: %s\n", hashedPassword)
    
    // Test correct password
    isCorrect := CheckPassword(password, hashedPassword)
    fmt.Printf("\nTesting correct password: %v\n", isCorrect)
    
    // Test wrong password
    wrongPassword := "wrongPassword123"
    isWrong := CheckPassword(wrongPassword, hashedPassword)
    fmt.Printf("Testing wrong password: %v\n", isWrong)
}
//github.com/gofiber/fiber/v2
