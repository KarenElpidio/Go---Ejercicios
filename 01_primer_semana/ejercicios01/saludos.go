// Ejercicio: En el main crear una función que reciba un nombre y devuelva un saludo con la forma “hola <nombre>”. Si el nombre recibido es vacío debe devolver “hola extraño”. Incluir un test para cada caso utilizando testify.

package main

import (
	"fmt"
)


func main(){
	var nombre string

	fmt.Printf("Coloca tu nombre: ")
	fmt.Scanf("%\n", &nombre)
	fmt.Printf("Hola %s\n", nombre)


	}