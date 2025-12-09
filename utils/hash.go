package utils

import "golang.org/x/crypto/bcrypt"

func HashData(data string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), 14)
	return string(bytes), err
}

func CheckDataHash(data, hashedData string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedData), []byte(data))
	return err == nil
	
}
