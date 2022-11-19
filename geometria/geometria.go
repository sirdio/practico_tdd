package geometria

/*
Dados los numeros de base y altura, cuando se llama a la funcion
CalcularArea, entonces retorna el Area de un Triangulo
*/
func Area(base, altura float64) interface{} {

	return (base * altura) / 2

}

/*
Dados tres lados de un triangulo, cuando se llama a la funcion
CalcularPerimetro, entonces retorna el Perimetro de un Triangulo
*/
func Perimetro( triangulo Triangulo) interface{} {
	return triangulo.Lado_1 + triangulo.Lado_2 + triangulo.Lado_3
}

// estructrua de tipo triangulo que contiene la longitudes de los lados del triangulo
type Triangulo struct{
	Lado_1 float64
	Lado_2 float64
	Lado_3 float64
}