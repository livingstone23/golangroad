package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	fmt.Println(paises)

	paises["mexico"] = "D.F."
	paises["españa"] = "madrid"
	paises["francia"] = "paris"

	fmt.Println(paises)
	fmt.Println(paises["españa"])

}

func MostrarMapas2() {
	campeonato := map[string]int{
		"Real Madrid": 39,
		"Chivas":      10,
		"Barcelona":   15,
		"Chealsea":    11,
	}

	fmt.Println(campeonato)

	//recorremos con el range
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, triene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
