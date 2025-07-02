package main

import (
	"fmt"
	"go-debt-book/middleware"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: debts [init | add-user <name> | add-debt <user_id> | get-book <user_id>]")
		return
	}

	switch args[0] {

	case "init":
		middleware.InitDB()

	case "add-user":
		if len(args) < 2 {
			fmt.Println("Missing full name")
			return
		}
		middleware.AddUser(args[1])

	case "add-debt":
		if len(args) < 3 {
			fmt.Println("Usage: go run main.go add-debt <user_id> <debt>")
			return
		}
		userID, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println("Invalid user ID")
			return
		}
		debtStr := args[2]
		debt, err := strconv.ParseFloat(debtStr, 64)
		if err != nil || debt < 0 {
			fmt.Println("Invalid debt amount (must be a non-negative number)")
			return
		}
		middleware.AddDebt(userID, debt)

	case "get-book":
		if len(args) < 2 {
			fmt.Println("Missing user ID")
			return
		}
		userID, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println("Invalid user ID")
			return
		}
		debts := middleware.GetBook(userID)
		for _, d := range debts {
			fmt.Printf("Debt ID: %d | Amount: %.2f | Created: %s\n", d.Debt_ID, d.Debt, d.Created_At.Format(time.RFC3339))
		}

	default:
		fmt.Println("Unknown command")
	}
}
