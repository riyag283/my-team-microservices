package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("my-secret-key")

func GenerateToken(userID int64, role string) (string, error) {
    // create a new token object
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId": userID,
        "role":   role,
        "exp":    time.Now().Add(time.Hour * 24).Unix(), // expires after 24 hours
    })

    // sign the token with our secret key
    signedToken, err := token.SignedString(secret)
    if err != nil {
        return "", err
    }

    return signedToken, nil
}
