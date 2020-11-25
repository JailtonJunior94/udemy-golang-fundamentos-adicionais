package enderecos

import (
	"strings"
)

// TipoDeEndereco - Verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoLetraMinuscula := strings.ToLower(endereco)

	primeiraPalavraEndereco := strings.Split(enderecoLetraMinuscula, " ")[0]
	enderecoTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTipoValido = true
		}
	}

	if enderecoTipoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo inválido"
}
