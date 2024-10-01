package main

import "fmt"

func main() {
	dia := diaDaSemana(6)
	fmt.Println(dia)

	mesTexto := mes(2)
	fmt.Println(mesTexto)

	estacaoAno := estacaoAno(6)
	fmt.Println(estacaoAno)
}

func diaDaSemana(numeroDia int) string {
	switch numeroDia {
		case 0:
			return "Domingo"
		case 1:
			return "Segunda-feira"
		case 2:
			return "Terça-feira"
		case 3:
			return "Quarta-feira"
		case 4:
			return "Quinta-feira"
		case 5:
			return "Sexta-feira"
		case 6:
			return "Sábado"
		default:
			return "Dia indefindo"
	}
}

func mes(numeroMes int) string {
	switch {
	case numeroMes == 1:
		return "Janeiro"
	case numeroMes == 2:
		return "Fevereiro"
	case numeroMes == 3:
		return "Março"
	case numeroMes == 4:
		return "Abril"
	case numeroMes == 5:
		return "Maio"
	case numeroMes == 6:
		return "Junho"
	case numeroMes == 7:
		return "Julho"
	case numeroMes == 8:
		return "Agosto"
	case numeroMes == 9:
		return "Setembro"
	case numeroMes == 10:
		return "Outubro"
	case numeroMes == 11:
		return "Novembro"
	case numeroMes == 12:
		return "Dezembro"
	default:
		return "Mês indefinido"
	}
}

func estacaoAno(numeroMes int) string {
	var estacao string

	switch {
	case numeroMes >= 1 && numeroMes <= 3:
		estacao = "Outono"
	case numeroMes >= 4 && numeroMes <= 6:
		estacao = "Inverno"
	case numeroMes >= 7 && numeroMes <= 9:
		estacao = "Primavera"
	case numeroMes >= 10 && numeroMes <= 12:
		estacao = "Verão"
	default:
		estacao = "Estação indefinida"
	}

	return estacao
}
