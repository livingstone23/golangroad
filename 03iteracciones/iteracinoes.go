package iteraciones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Iterar() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func Iterar2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func Iterar3() {
	for i := 0; i <= 5; i++ {

		if i == 3 {
			break
		}

		fmt.Println(i)
	}
}

func Iterar4() {
	for i := 0; i <= 5; i++ {

		if i == 3 {
			continue
		}

		fmt.Println(i)
	}
}

func TablaDeMultiplicar() string {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error
	var texto string

	for {
		fmt.Println("Ingrese un numero : ")
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
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}
	//fmt.Println(texto)
	return texto

}
