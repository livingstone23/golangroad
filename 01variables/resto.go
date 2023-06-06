package variables

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float64
var Fecha time.Time

func RestoVariables() {
	Nombre = "Livingstone"
	Estado = true
	Sueldo = 33000
	Fecha = time.Now()
	fmt.Println("El valor de la variable Nombre es: ", Nombre)
	fmt.Println("El valor de la variable Estado es: ", Estado)
	fmt.Println("El valor de la variable Sueldo es: ", Sueldo)
	fmt.Println("El valor de la variable Fecha es: ", Fecha)

}

func ConviertoTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)

	return true, texto
}

func IdentificarSistema() string {

	os := runtime.GOOS
	/*


		if os == "Linux." || os == "OS X." {
			fmt.Println("Esto no es windows, es: ", os)
		} else {
			fmt.Println("Esto es windows", os)
		}

	*/

	switch os2 := runtime.GOOS; os2 {
	case "linux":
		fmt.Println("Esto es linux")
		return os2
	case "darwin":
		fmt.Println("Esto es darwin")
		return os2
	default:
		fmt.Println("%s \n", os2)

	}

	return os
}
