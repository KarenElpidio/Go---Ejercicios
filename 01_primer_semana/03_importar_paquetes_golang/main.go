package main 

import (
	// Paquete de golang, no necesitamos manejo de dependencias
	//basta con importarlo asi 

	"fmt"
	"math/rand" 
	"time"
)

func main ()  {
	//Llamamos u una función detro del mismo paquete, no necesita ser exportada
	imprimirNumeroRand()
}

// ImprimirNumeroRand utiliza el paquete rand para imprimir número.

func imprimirNumeroRand() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Mi número favorito es %d\n", rand.Intn(10))
}
