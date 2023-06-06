package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {

	//os.Stdin es Standar Input y el standar input es el teclado
	//os.Stdout es Standar Outuput y es la pantalla.
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 1: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			//Para abortar la aplicacion
			panic("El dato ingresado es incorrecto : " + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			//Para abortar la aplicacion
			panic("El dato ingresado es incorrecto : " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda : ")
	if scanner.Scan() {
		leyenda = scanner.Text()

	}

	fmt.Println(leyenda, numero1*numero2)
}

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error

	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("El valor de %d*%d es = %d\n ", numero, i, i*numero)
	}

}
