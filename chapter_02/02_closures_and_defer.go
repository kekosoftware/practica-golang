package main

import (
	"fmt"
	"os"
	"log"
)

var nextNumber int = 0

// closures
func sequenceGenerator() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// defer
func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	// Process file
	fmt.Println("Se va a procesar el archivo")
}

func main() {
	nextNumber := sequenceGenerator()
	fmt.Println(nextNumber()) // Output: 1
	fmt.Println(nextNumber()) // Output: 2
	// Llamar a la función readFile con el nombre del archivo de prueba
	filename := "archivo.txt"
	readFile(filename)
}
