package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	//vivo       bool
}

// Para implementar la interface
// Automaticamente  go realiza que aplica
// interface porque implementa las mismas funcionalidades
func (h *Hombre) Respirar()    { h.respirando = true }
func (h *Hombre) Pensar()      { h.pensando = true }
func (h *Hombre) Comer()       { h.comiendo = true }
func (h *Hombre) Sexo() string { return "Hombre" }
