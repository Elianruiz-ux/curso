package main

import "fmt"

func main(){
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Juan"
	edad := 20

	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, %s y su edad es %d \n", nombre, edad)

	fmt.Println(mensaje)

	fmt.Printf("nombre: %T \n", nombre)

	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("El otro nombre es: ", nombre)
}