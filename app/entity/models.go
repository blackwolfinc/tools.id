package entity

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Address   string
	Role      string
	CreatedAt time.Time
}
