package utils_test

import (
	"testing"

	"github.com/hack-caixa/domain"
	"github.com/hack-caixa/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestMakePriceParcelasDTOFunction(t *testing.T) {
	items := make([]domain.ParcelaDTO, 5)

	input := domain.NewInput()
	input.Prazo = 5
	input.ValorDesejado = 900.00

	produto := domain.NewProduct()
	produto.CO_PRODUTO = 1
	produto.NO_PRODUTO = "Produto 1"
	produto.NU_MAXIMO_MESES = 24
	produto.NU_MINIMO_MESES = 0
	produto.PC_TAXA_JUROS = 0.017900000
	produto.VR_MINIMO = 200.00
	produto.VR_MAXIMO = 10000.00

	value := utils.MakePriceParcelasDTO(items, *input, *produto)

	items = []domain.ParcelaDTO{{Numero: 1, ValorAmortizacao: 173.67, ValorJuros: 16.11, ValorPrestacao: 189.78}, {Numero: 2, ValorAmortizacao: 176.78, ValorJuros: 13, ValorPrestacao: 189.78}, {Numero: 3, ValorAmortizacao: 179.94, ValorJuros: 9.84, ValorPrestacao: 189.78}, {Numero: 4, ValorAmortizacao: 183.16, ValorJuros: 6.62, ValorPrestacao: 189.78}, {Numero: 5, ValorAmortizacao: 186.44, ValorJuros: 3.34, ValorPrestacao: 189.78}}
	resp := &domain.ResultadoSimulacaoDTO{Tipo: "PRICE", Parcelas: &items}

	require.Equal(t, value, resp)

}
