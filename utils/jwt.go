package utils

import (
	"log"
	"time"

	userModel "employee/model"

	"github.com/dgrijalva/jwt-go"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("secret")
)

// GenerateToken generates a jwt token and assign a useObj to it's claims and return it
func GenerateToken(useObj userModel.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["useObj"] = useObj
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the useObj in it's claims
func ParseToken(tokenStr string) (userModel.User, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	useObj := new(userModel.User)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tmpObj := claims["useObj"].(map[string]interface{})
		useObj.Email = tmpObj["email"].(string)
		useObj.Password = tmpObj["password"].(string)

		return *useObj, nil
	} else {
		return *useObj, err
	}
}
