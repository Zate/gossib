package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	// "log"
)

var CHIPER_KEY = []byte("0987654321234567")

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

func StrEncrypt(key []byte, message string) (encmess string, err error) {
	plain := []byte(message)

	block, err := aes.NewCipher(key)
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

func StrDecrypt(key []byte, securemess string) (decodedmess string, err error) {
	cText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
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
