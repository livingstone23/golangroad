package main

import (
	"fmt"
	//dar "github.com/livingstone23/golangroad/01variables"
	//tec "github.com/livingstone23/golangroad/02teclado"
	//ite "github.com/livingstone23/golangroad/03iteracciones"
	//arc "github.com/livingstone23/golangroad/04files"
	//fun "github.com/livingstone23/golangroad/05funciones"
	//arr "github.com/livingstone23/golangroad/06arreglos_slices"
	//mapas "github.com/livingstone23/golangroad/07mapas"
	//oob "github.com/livingstone23/golangroad/08orientacion_objetos"
	intf "github.com/livingstone23/golangroad/09interfaces"
)

func main() {
	fmt.Println("Iniciando el camino a Golang")

	//Leccion 01 Inicio con variables
	/*
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
	*/

	//Leccion 02 - Teclado
	//tec.IngresoNumeros()
	//tec.TablaMultiplicar()

	//Leccion 03 - Iteracciones
	//ite.Iterar()
	//ite.Iterar2()
	//ite.Iterar3()
	//ite.Iterar4()
	//fmt.Println(ite.TablaDeMultiplicar())

	//Leccion 04 - Archivos
	//arc.GrabarTabla()
	//arc.SumaTabla()
	//arc.LeoArchivo()
	//arc.LeoArchivo2()

	//Leccion 05 - Funciones
	//fun.Calculos()
	//fun.LlamarClosure()
	//fun.Exponencia(2)

	//Lecciones 06 - Arreglos
	//arr.MuestroArreglos()
	//arr.MuestroSlice()
	//arr.Capacidad()

	//Leccion 07 - Mapas
	//mapas.MostrarMapas()
	//mapas.MostrarMapas2()

	//Leccion 08 - Orientacion Objetos
	//oob.AltaUsuario()

	//Leccion 09 - Interfaces
	Livingstone := new(intf.Hombre)
	intf.HumanoRespirando(Livingstone)

	Mjose := new(intf.Mujer)
	intf.HumanoRespirando(Mjose)

}
