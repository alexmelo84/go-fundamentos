package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	endereco := "Alameda Barros"
	tipoDeEnderecoEsperado := "Alameda"
	tipoDeEnderecoRecebido := TipoDeEndereco(endereco)

	if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
		t.Errorf(
			"Tipo de endereço diferente do esperado! Esperava %s e recebeu %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido,
		)
	}
}
