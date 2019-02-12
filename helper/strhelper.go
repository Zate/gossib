/*
MIT License

Copyright (c) 2018 Aris Ripandi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

var cipherKey = []byte("0987654321234567")

func init() {
	assertAvailablePRNG()
}

/*
func main() {
    msg := "A quick brown fox jumped over the lazy dog."
	if encrypted, err := StrEncrypt(CHIPER_KEY, msg); err != nil {
		log.Println(err)
	} else {
		log.Printf("ENCRYPTED: %s\n", encrypted)
		if decrypted, err := StrDecrypt(CHIPER_KEY, encrypted); err != nil {
			log.Println(err)
		} else {
			log.Printf("DECRYPTED: %s\n", decrypted)
		}
	}
}
*/

func StrEncrypt(message string) (encmess string, err error) {
	plain := []byte(message)

	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return
	}

	cText := make([]byte, aes.BlockSize+len(plain))
	iv := cText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cText[aes.BlockSize:], plain)

	//returns to base64 encoded string
	encmess = base64.URLEncoding.EncodeToString(cText)
	return
}

func StrDecrypt(securemess string) (decodedmess string, err error) {
	cText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return
	}

	if len(cText) < aes.BlockSize {
		err = errors.New("cText block size is too short!")
		return
	}

	iv := cText[:aes.BlockSize]
	cText = cText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cText, cText)

	decodedmess = string(cText)
	return
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func StrRand(n ...int) (string, error) {
	lp := 8
	if len(n) > 0 {
		lp = n[0]
	}
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes, err := GenerateRandomBytes(lp)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

// GenerateRandomStringURLSafe returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func StrRandURLSafe(n ...int) (string, error) {
	lp := 8
	if len(n) > 0 {
		lp = n[0]
	}
	b, err := GenerateRandomBytes(lp)
	return base64.URLEncoding.EncodeToString(b), err
}
