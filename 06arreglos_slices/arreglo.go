package arreglo

import "fmt"

var tabla [10]int = [10]int{10, 3, 78}
var matriz [4][5]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[9] = 54

	fmt.Println(tabla)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[2][1] = 15
	fmt.Println(matriz)
}
