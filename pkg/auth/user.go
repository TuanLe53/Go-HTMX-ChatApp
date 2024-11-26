package auth

import "golang.org/x/crypto/bcrypt"

func HashPw(pw string) (string, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(hashedPw), err
}

func CheckPw(hashedPw, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
}
