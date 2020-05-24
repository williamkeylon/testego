package goarea

import "math"

//Pi Numero do Pi relação entre perimetro e uma circunferencia
const Pi = 3.1415

//Circ calcula o circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect caucula retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
