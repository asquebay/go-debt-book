package models

import "time"

// User represents a user in the system
type User struct {
	User_ID  int64  `json:"User_ID"`
	FullName string `json:"FullName"`
}

// Debt represents a debt associated with a user
type Debt struct {
	Debt_ID    int64     `json:"Debt_ID"`
	User_ID    int64     `json:"User_ID"`
	Debt       float64   `json:"Debt"`
	Created_At time.Time `json:"Created_At"`
}
