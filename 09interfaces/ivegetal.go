package interfaces

type Vegetal interface {
	ClasificacionVegetal() string
	EstaVivo() bool //por tener esta propiedad califica instancia interface SerVivo
}
