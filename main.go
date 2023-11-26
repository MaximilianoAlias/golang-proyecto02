package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/MaximilianoAlias/golang-proyecto02/filecrypt"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		imprimirAyuda()

		os.Exit(0)
	}
	//LLAMAMOS A LAS 3 FUNCIONES PRINCIPALES DESDE NUESTRA FUNCION PRINCIPAL
	function := os.Args[1]

	switch function {
	case "ayuda":
		imprimirAyuda()

	case "encriptar":
		encryptHandle()

	case "desencriptar":
		decryptHandle()

	default:
		fmt.Println("Inicia la funcion encriptar para encriptar y la funcion desencriptar para desencriptar un archivo")
		os.Exit(1)
	}
}

func imprimirAyuda() {
	fmt.Println("encriptación de archivos")
	fmt.Println("Cifrador de archivos simple para tus necesidades diarias.")
	fmt.Println("")
	fmt.Println("Uso:")
	fmt.Println("")
	fmt.Println("\tgo run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Comandos:")
	fmt.Println("")
	fmt.Println("\t encriptar\tEncripta un archivo otorgando una contraseña")
	fmt.Println("\t desencriptar\tIntenta descifrar un archivo usando una contraseña")
	fmt.Println("\t ayuda\t\tMuestra texto de ayuda")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		println("Se necesita la ruta del archivo, para más información ejecute . ayuda")
		os.Exit(0)
	}

	file := os.Args[2]
	if !validarArchivo(file) {
		panic("El archivo no existe.")
	}

	contraseña := obtenerContraseña()
	fmt.Println("\nEncriptando...")
	// LLAMO AL ARCHIVO filecrypt.go y al método encriptar

	filecrypt.Encriptar(file, contraseña)
	fmt.Println("\nArchivo encriptado exitosamente")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		println("Se necesita la ruta del archivo, para mas informacion ejecute . ayuda")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validarArchivo(file) {
		panic("El archivo no existe.")
	}

	fmt.Println("Ingrese la contraseña para desencriptar el archivo.")
	contraseña, _ := term.ReadPassword(0)
	fmt.Println("\nDesencriptando...")

	filecrypt.Desencriptar(file, contraseña)
	fmt.Println("\nArchivo desencriptado exitosamente")
}

func obtenerContraseña() []byte {
	fmt.Println("Ingrese la contraseña")
	contraseña, _ := term.ReadPassword(0)
	fmt.Println("\n Confirmar la contraseña")
	contraseña2, _ := term.ReadPassword(0)

	if !validatePassword(contraseña, contraseña2) {
		fmt.Println("\n Las contraseñas no coinciden, intente nuevamente")
		return obtenerContraseña()
	}
	return contraseña
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}

	return true
}

func validarArchivo(file string) bool {

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true

}
