package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()
	//auxiliar.escrever2() aqui daria erro pois a função escrever não é publica

	erro := checkmail.ValidateFormat("123")
	fmt.Println((erro))

}
