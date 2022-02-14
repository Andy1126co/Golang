package main

import "fmt"

func main() {
	//
	//// Area cuadrado
	//const baseCuadrado = 10
	//areaCuadrado := baseCuadrado * baseCuadrado
	//fmt.Println("Area cuadrado: ", areaCuadrado)
	//
	//x := 10
	//y := 50
	//
	////suma
	//result := x + y
	//fmt.Println("suma: ", result)
	//
	////Resta
	//result = y - x
	//fmt.Println("Resta: ", result)
	//
	////multiplicacion
	//result = x * y
	//fmt.Println("multiplicacion: ", result)
	//
	////división
	//result = y / x
	//fmt.Println("Division: ", result)
	//
	////modulo
	//result = y % x
	//fmt.Println("Modulo: ", result)
	//
	////Incrementar
	//x++
	//fmt.Println("Incrementa", x)
	//
	////Decremental
	//x--
	//fmt.Println("Decremental: ", x)
	//
	////Retos
	////area Rectangulo,trapecio, circulo
	//
	////Area rectangulo
	//b := 10
	//a := 23
	//h := 12
	//
	//areaRectangulo := b * a
	//fmt.Println("Area del rectangulo: ", areaRectangulo)
	//
	////Area trapecio
	//areaTrapecio := ((b + a) * h) / 2
	//fmt.Println("Area del trapecio: ", areaTrapecio)
	//
	////area circulo
	//pi := 3.1415
	//var radio float64 = 5
	//areaCirculo := pi * (math.Pow(radio, 2))
	//fmt.Println("Area del circulo: ", areaCirculo)
	//Println

	// Declaración de variables
	helloMessage := "hello"
	worldMessage := "world"

	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Daniel"
	cursos := 2
	fmt.Printf("%s a finalizado %d cursos\n", nombre, cursos)
	fmt.Printf("%v a finalizado %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s a finalizado %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo Datos
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("worldMessage: %T \n", cursos)
}
