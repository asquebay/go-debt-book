package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"go-debt-book/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// createConnection establishes a connection to PostgreSQL
func createConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

// InitDB creates the necessary tables if they don't exist
func InitDB() {
	db := createConnection()
	defer db.Close()

	// Create Users table
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            User_ID SERIAL PRIMARY KEY,
            FullName VARCHAR(45)
        )`)
	if err != nil {
		log.Fatal(err)
	}

	// Create Debts table
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS Debts (
        Debt_ID SERIAL PRIMARY KEY,
        User_ID INT REFERENCES Users(User_ID),
        Debt NUMERIC(15,2),
        Created_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tables initialized successfully.")
}

// AddUser inserts a new user into the Users table
func AddUser(fullName string) int64 {
	db := createConnection()
	defer db.Close()

	var id int64
	err := db.QueryRow("INSERT INTO Users (FullName) VALUES ($1) RETURNING User_ID", fullName).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User %s added with ID: %d\n", fullName, id)
	return id
}

// AddDebt creates a new debt entry for a given user
func AddDebt(userID int64, amount float64) {
	db := createConnection()
	defer db.Close()

	_, err := db.Exec("INSERT INTO Debts (User_ID, Debt) VALUES ($1, $2)", userID, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Debt %.2f added for user %d\n", amount, userID)
}

// GetBook fetches all debts for a given user
func GetBook(userID int64) []models.Debt {
	db := createConnection()
	defer db.Close()

	rows, err := db.Query("SELECT Debt_ID, Debt, Created_At FROM Debts WHERE User_ID = $1", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var debts []models.Debt
	for rows.Next() {
		var debt models.Debt
		err := rows.Scan(&debt.Debt_ID, &debt.Debt, &debt.Created_At)
		if err != nil {
			log.Fatal(err)
		}
		debt.User_ID = userID
		debts = append(debts, debt)
	}
	return debts
}
