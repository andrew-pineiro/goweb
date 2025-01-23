package controllers

import (
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(user User) string {
	hasher := sha256.New()
	hasher.Write([]byte(user.Username))
	hasher.Write([]byte(string(user.Id)))
	hasher.Write([]byte(user.PasswordHash))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
