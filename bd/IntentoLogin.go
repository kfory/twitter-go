package bd

import (
	"twitter-go/models"

	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usu, encontrado, _ := ChequeoYaExisteUsuario(email)

	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBd := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBd, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
