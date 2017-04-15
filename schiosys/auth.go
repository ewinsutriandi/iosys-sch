package schiosys

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

func generateRandomSalt(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	HandleError(err)
	return b, nil
}

func encrypt(pass string, salt []byte) ([]byte, error) {
	dk, err := scrypt.Key([]byte(pass), salt, 16384, 8, 1, 32)
	return dk, err
}

//Authenticate user using username and password
func Authenticate(username, pass string) bool {
	//savedPass = get []byte from db
	savedPass, salt, err := getPassSalt(username)
	if err == nil {
		key, err := encrypt(pass, salt)
		if err == nil {
			i := bytes.Compare(savedPass, key)
			if i == 0 {
				return true
			}
		}
	}
	HandleError(err)
	return false
}

//get pass & salt from db for user username
func getPassSalt(username string) ([]byte, []byte, error) {
	var err error
	return nil, nil, err
}
