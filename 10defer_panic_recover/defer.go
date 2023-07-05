package defer_panic

import (
	"fmt"
	"log"
	"time"
)

func DemosDefer() {
	//Defer permite controlar ultima instruccion en ejecutarse

	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es el segundo mensaje")
}

// EjemploPanic muestra la funcionalidad de panic
// se agrega a la funcion () para volverla una funcion anonima
func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrió un error que generó Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}

func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func ConectarDB() string {
	user := "user"
	password := "password"
	host := "host"
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	return connStr

}
