package main

import "fmt"

const ebulicaoK = 373

func main(){

	c := converterKevinToCelsius(ebulicaoK)

	fmt.Println("O valor de ebulição em Kelvin é",ebulicaoK, "e em celsius é: ",c)
}

func converterKevinToCelsius(k int) int{
	return k - 273
}