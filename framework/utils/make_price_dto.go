package utils

import (
	"fmt"
	"strconv"

	"github.com/hack-caixa/domain"
)

func MakePriceParcelasDTO(price []domain.ParcelaDTO, in domain.EntradaSimulacaoDTO, produto domain.PRODUTO) *domain.ResultadoSimulacaoDTO {
	for i := range price {
		listOfAmortization := CalcularAmortizacaoPrice(float64(in.ValorDesejado), float64(produto.PC_TAXA_JUROS), int(in.Prazo))
		var amortizacao float32
		if i != 0 {
			amortizacao = listOfAmortization[i] - listOfAmortization[i-1]
		} else {
			amortizacao = listOfAmortization[i]
		}
		amt := fmt.Sprintf("%.2f", amortizacao)
		amort, _ := strconv.ParseFloat(amt, 32)

		prestacao := float32(CalcularParcelaPrice(float64(in.ValorDesejado), float64(produto.PC_TAXA_JUROS), int(in.Prazo)))

		juros := float32(CalcularJurosPrice(float64(in.ValorDesejado), float64(produto.PC_TAXA_JUROS), int(in.Prazo), i+1, listOfAmortization))

		price[i] = domain.ParcelaDTO{
			Numero:           int32(i + 1),
			ValorAmortizacao: float32(amort),
			ValorJuros:       float32(juros),
			ValorPrestacao:   float32(prestacao),
		}

	}

	resultadoPrice := &domain.ResultadoSimulacaoDTO{Tipo: "PRICE", Parcelas: &price}

	return resultadoPrice
}
