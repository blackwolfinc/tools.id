package auth

import "fmt"

func HashPassword(password string) string {
	// Password hashing logic
	return fmt.Sprintf("hashed_%s", password)
}
