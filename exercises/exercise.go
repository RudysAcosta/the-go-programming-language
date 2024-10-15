package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	// createArchivo("test.txt", "Hola mundo Cruel")
	readArchivo("test.txt")
	readWorld()
}

func createArchivo(name string, text string) (*bufio.ReadWriter, error) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("%s\n", text))
	if err != nil {
		return nil, err
	}

	err = writer.Flush()
	if err != nil {
		return nil, err
	}

	return bufio.NewReadWriter(bufio.NewReader(file), writer), nil
}

func readArchivo(name string) {

	file, err := os.Open(name)

	if err != nil {
		fmt.Println("Error al leer el archivo")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 1

	for scanner.Scan() {
		fmt.Printf("%d) %s\n", i, scanner.Text())
		i += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error durante el escaneo:", err)
	}

}

func readWorld() {
	text := "Hola mundo, esto es un ejemplo"

	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	i := 1
	for scanner.Scan() {
		fmt.Printf("%d) %s \n", i, scanner.Text())
	}
}
