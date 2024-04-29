package main

import "fmt"

func main(){
	//arrays
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)
	fmt.Println(numeros[1])

	// Arrays con valores
	nombres := [3]string{"Elian", "Joel", "Ruiz"}
	fmt.Println(nombres)

	colores := [...]string{
		"Rojos",
		"Verde",
		"Negro",
		"Azul",
		"Amarillo",
	}
	fmt.Println(colores, len(colores))

	// Indices definidos
	monedas := [...]string{
		0:"Dolar",
		2:"Soles",
		3:"Pesos",
		5:"Euro",
	}
	fmt.Println(monedas, len(monedas))

	// Sub arrays

	subMoneda := monedas[3:2]
	fmt.Println(subMoneda)
}
