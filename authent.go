package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string
	PasswordHash string
}

var users = map[string]User{}

func register(username, password string) error {
	if _, exists := users[username]; exists {
		return fmt.Errorf("username already taken")
	}

	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Save the user with the hashed password
	users[username] = User{
		Username:     username,
		PasswordHash: string(passwordHash),
	}

	return nil
}

func login(username, password string) error {
	user, exists := users[username]
	if !exists {
		return fmt.Errorf("user not found")
	}

	// Compare the stored hashed password with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Print("Enter username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Print("Enter password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			err := register(username, password)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Registration successful")
			}
		case "2":
			fmt.Print("Enter username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Print("Enter password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			err := login(username, password)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Login successful")
			}
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
