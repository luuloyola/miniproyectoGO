package main

import "fmt"

func main() {
	// Para saludar al usuario y pedir los puntajes
	fmt.Println("Bienvenido a la encuesta, a continuacion se le pediran 10 puntajes, con valores validos entre 1 y 5: ")
	// Arreglo que almacenara la frecuencia de puntajes
	arrPuntajes := []int{0, 0, 0, 0, 0}
	var x int
	// for
	for i := 0; i < 10; i++ {
		fmt.Printf("Ingrese el puntaje numero %d \n", i+1)
		fmt.Scan(&x)
		for x < 1 || x > 5 {
			fmt.Print("Debe ingresar un puntaje entre 1 y 5:\n")
			fmt.Scan(&x)
		}
		x -= 1
		arrPuntajes[x] += 1
	}

	for i := 0; i < len(arrPuntajes); i++ {
		fmt.Printf("\nEl puntaje %d aparece %d ", i+1, arrPuntajes[i])
		if arrPuntajes[i] == 1 {
			fmt.Printf("vez\n")
		} else {
			fmt.Printf("veces\n")
		}
	}

	if arrPuntajes[0]+arrPuntajes[1] < arrPuntajes[3]+arrPuntajes[4] {
		fmt.Print("\n\t¡Buen resultado!")
	} else {
		fmt.Print("\n\tMejora.")
	}
}
