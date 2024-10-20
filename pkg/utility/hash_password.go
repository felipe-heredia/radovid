package utility

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, _ := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes)
}

func CompareHasAndPassword(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}
