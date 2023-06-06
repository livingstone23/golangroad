package main

import (
	"fmt"
	dar "github.com/livingstone23/golangroad/01variables"
)

func main() {
	fmt.Println("Iniciando el camino a Golang")

	//Leccion 01 Inicio con variables
	dar.MostrarEnteros()
	dar.RestoVariables()
	estado, resultado := dar.ConviertoTexto(100)
	if estado != false {
		fmt.Printf("El resultado de la conversion es : %v", resultado)
	} else {
		fmt.Println("no se pudo realizar la conversion")
	}
	sistema := dar.IdentificarSistema()
	fmt.Println("El sistema operativo es: ", sistema)

}
