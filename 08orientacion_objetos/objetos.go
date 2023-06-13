package modelos

import (
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(User)
	u.AddUser(10, "Livingstone", time.Now(), true)
	fmt.Println(u)
}
