package goarea

import "math"

// Pi é uma proporção numérica definids pela relação entre
// o perímetro de uma circunfer~encia e seu diâmetro
const Pi = 3.1416

func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func Rect(base, altura float64) float64 {
	return base * altura
}

func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
