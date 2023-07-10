package main

import "fmt"

const ebulicaoK float64 = 373.0
const fusaoK float64 = 273.0

func main() {

	tempC := (ebulicaoK - fusaoK) // ebulicao de K para  C
	fmt.Println("A temperatura de Ebulicao da agua em K = ", ebulicaoK)
	fmt.Println("A temperatura de Fusao da agua em K = ", fusaoK)
	fmt.Printf("Conversao da temperatura de ebulicao da agua de K para C = %v  ", tempC)
}
