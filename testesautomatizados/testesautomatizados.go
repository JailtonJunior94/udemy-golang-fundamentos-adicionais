package main

import (
	"fmt"

	"github.com/jailtonjunior94/udemy-golang-fundamentos-adicionais/testesautomatizados/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
