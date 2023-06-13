package interfaces

type Animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	EstaVivo() bool //por tener esta propiedad califica instancia interface SerVivo
}
