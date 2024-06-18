package pass

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

const (
	randChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func Random() (password string, hashedpassword string, err error) {
	randPassword, errRandPass := randCharacter(10)
	if errRandPass != nil {
		return "", "", errRandPass
	}
	hashPass, errHashPass := HashPassword(randPassword)
	if errHashPass != nil {
		return "", "", errHashPass
	}
	return randPassword, string(hashPass), nil
}

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func ComparePassword(userPassword []byte, dtoPassword string) error {
	err := bcrypt.CompareHashAndPassword(userPassword, []byte(dtoPassword))
	return err
}

func randCharacter(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	charsetLength := len(randChar)
	for i := 0; i < length; i++ {
		buffer[i] = randChar[int(buffer[i])%charsetLength]
	}
	return string(buffer), nil
}

func RandUlid() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	rand, _ := ulid.New(ms, entropy)
	return rand.String()
}
