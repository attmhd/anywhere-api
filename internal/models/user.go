package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Fullname  string    `json:"fullname" db:"fullname"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Phone     string    `json:"phone" db:"phone"`
	Country   string    `json:"country" db:"country"`
	Gender    bool      `json:"gender" db:"gender"`
	Address   string    `json:"address" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
