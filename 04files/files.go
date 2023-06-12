package files

import (
	"bufio"
	"fmt"
	ite "github.com/livingstone23/golangroad/03iteracciones"
	"io/ioutil"
	"os"
)

var fileName string = "./04files/txt/tabla.txt"

func GrabarTabla() {
	var texto string = ite.TablaDeMultiplicar()
	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
	}
	fmt.Fprintln(archivo, texto)
	//Todo archivo que se abre se cierra
	archivo.Close()

}

func SumaTabla() {
	var texto string = ite.TablaDeMultiplicar()

	if Append(fileName, texto) == false {
		fmt.Println("Error al concatenar contenio")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	fmt.Println(string(archivo))
}

func LeoArchivo2() {
	archivo, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	//Se debe cerrar el archivo
	archivo.Close()
}
