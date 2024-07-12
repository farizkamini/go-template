package pass

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

func Random() (password string, hashedpassword string, err error) {
	randPassword := generateRandomString(10)
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

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	seededRand := rand.New(source)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(result)
}

func RandUlid() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	rand, _ := ulid.New(ms, entropy)
	return rand.String()
}
