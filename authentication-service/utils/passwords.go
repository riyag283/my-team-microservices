package utils

import "golang.org/x/crypto/bcrypt"

const cost int = 10

// EncryptPassword returns the bcrypt hash of the given password
func EncryptPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// VerifyPassword compares a plaintext password with its corresponding hashed password
func VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
