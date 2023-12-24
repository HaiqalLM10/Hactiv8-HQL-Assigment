package utility

import "golang.org/x/crypto/bcrypt"

func BcryptHash(plainText string) (string, error) {
	byteHash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(byteHash), nil
}
