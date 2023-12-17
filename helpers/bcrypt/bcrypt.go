package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(inputPassword, password string) error {
	hashedPasswordFromDatabase := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashedPasswordFromDatabase, []byte(inputPassword))
	if err != nil {
		return err
	}
	return nil
}
