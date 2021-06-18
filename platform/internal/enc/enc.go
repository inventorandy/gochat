package enc

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

// Random Charset
const randCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890!@$%^&*()-_{}[]:;,.?~"

// GenerateSalt generates a randomised string of length {length}
func GenerateSalt(length int) string {
	// Random Seed Generator
	rand.Seed(time.Now().UnixNano())

	// Create the Byte Array
	b := make([]byte, length)

	// Populate the Byte Array
	for i := range b {
		b[i] = randCharset[rand.Intn(len(randCharset))]
	}

	// Return the String
	return string(b)
}

// SHA1EncryptWithSalt encrypts a string {input} with a salt {salt} appended
func SHA1EncryptWithSalt(input string, salt string) string {
	// Create the Hash
	hash := sha1.New()

	// Concat the Input and Salt
	encodableString := fmt.Sprintf("%s%s", input, salt)

	// Hash the String
	hash.Write([]byte(encodableString))
	bs := hash.Sum(nil)

	// Return the String
	return fmt.Sprintf("%x", bs)
}

// SHA256EncryptWithSalt encrypts a string {input} with a salt {salt} appended
func SHA256EncryptWithSalt(input string, salt string) string {
	// Create the Hash
	hash := sha256.New()

	// Concat the Input and Salt
	encodableString := fmt.Sprintf("%s%s", input, salt)

	// Hash the String
	hash.Write([]byte(encodableString))
	bs := hash.Sum(nil)

	// Return the String
	return fmt.Sprintf("%x", bs)
}
