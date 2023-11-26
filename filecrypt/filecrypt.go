package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encriptar(source string, password []byte) {
	//ANALIZO LA RUTA Y EL ARCHIVO EXISTE
	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	//SI EXISTE EL ARCHIVO LO ABRE
	srcFile, err := os.Open(source)

	if err != nil {
		panic(err.Error())
	}

	defer srcFile.Close()

	//TEXTO SIN FORMATO
	plaintext, err := io.ReadAll(srcFile)

	if err != nil {
		panic(err.Error())
	}

	key := password

	nonce := make([]byte, 12)

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//CREACION DE CLAVE DERIVADA
	//PARAMETROS KEY, NONCE, CANTIDAD DE ITERACIONES, LEE EL LOS ULTIMOS 12 DIGITOS DEL ARCHIVO CIFRADO
	//USO DE PASSWORD-BASED KEY DERIVATION FUNCTION
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)

	if err != nil {
		panic(err.Error())
	}

	aesgmc, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgmc.Seal(nil, nonce, plaintext, nil)

	ciphertext = append(ciphertext, nonce...)

	//DESTINATION FILE
	dstFile, err := os.Create(source)

	if err != nil {
		panic(err.Error())
	}

	defer dstFile.Close()

	_, err = dstFile.Write(ciphertext)

	if err != nil {
		panic(err.Error())
	}

}

func Desencriptar(source string, password []byte) {

	if _, err := os.Stat(source); os.IsNotExist(err) {
		panic(err.Error())
	}

	//SI EXISTE EL ARCHIVO LO ABRE
	srcFile, err := os.Open(source)

	if err != nil {
		panic(err.Error())
	}

	defer srcFile.Close()

	//PASAMOS EL ARCHIVO CIFRADO
	ciphertext, err := io.ReadAll(srcFile)

	if err != nil {
		panic(err.Error())
	}

	key := password

	//CREAMOS UNA VARIABLE QUE ALMACENA EL CIFRADO Y LA LONGITUD SIN LOS ULTIMOS 12
	salt := ciphertext[len(ciphertext)-12:]

	str := hex.EncodeToString(salt)

	nonce, err := hex.DecodeString(str)

	//CREACION DE CLAVE DERIVADA
	dk := pbkdf2.Key(key, nonce, 4096, 32, sha1.New)

	block, err := aes.NewCipher(dk)

	if err != nil {
		panic(err.Error())
	}

	aesgmc, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgmc.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)

	if err != nil {
		panic(err.Error())
	}

	//DESTINATION FILE
	dstFile, err := os.Create(source)

	if err != nil {
		panic(err.Error())
	}

	defer dstFile.Close()

	_, err = dstFile.Write(plaintext)

	if err != nil {
		panic(err.Error())
	}
}
