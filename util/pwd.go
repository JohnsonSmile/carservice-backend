package util

import "golang.org/x/crypto/bcrypt"

func GeneratePasswordHash(pwd string) (string, error) {
	cryptBytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(cryptBytes), err
}

func VerifyPasswordHash(hashedPwd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	return err == nil
}
