package utils

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "io"
)

// Encrypt encrypts sensitive data using AES encryption
func Encrypt(text string) (string, error) {
    key := []byte(os.Getenv("ENCRYPTION_KEY"))
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    ciphertext := make([]byte, aes.BlockSize+len(text))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))

    return base64.URLEncoding.EncodeToString(ciphertext), nil
}