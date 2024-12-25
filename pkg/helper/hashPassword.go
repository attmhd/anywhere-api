package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Generate the hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "Failed Hash Password", err
	}
	return string(hashedPassword), nil
}
