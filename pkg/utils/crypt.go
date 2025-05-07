package utils

import (
	"github.com/alf-grindel/dawn/pkg/constants"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	addSalt := password + " " + constants.SALT
	b, err := bcrypt.GenerateFromPassword([]byte(addSalt), bcrypt.DefaultCost)
	if err != nil {
		log.Println("utils.crypt: generate password failed")
		return ""
	}
	return string(b)
}

func ComparePassword(password, hashPassword string) bool {
	addSalt := password + " " + constants.SALT
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(addSalt))
	if err != nil {
		log.Println("utils.crypt: compare password failed")
		return false
	}
	return true
}
