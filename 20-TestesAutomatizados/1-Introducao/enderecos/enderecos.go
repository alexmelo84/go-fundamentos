package enderecos

import(
	"strings"
)

// TipoDeEndereco valida e retorna o tipo de endereço
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "alameda", "praça"}

	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoMinusculo, " ")[0]

	enderecoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo inválido"
}
