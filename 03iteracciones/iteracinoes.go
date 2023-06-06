package iteraciones

import "fmt"

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
