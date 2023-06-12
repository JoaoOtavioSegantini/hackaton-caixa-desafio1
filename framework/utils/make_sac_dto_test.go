package utils_test

import (
	"testing"

	"github.com/hack-caixa/domain"
	"github.com/hack-caixa/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestMakeSACParcelasDTOFunction(t *testing.T) {
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

	value := utils.MakeSacParcelasDTO(items, *input, *produto)

	items = []domain.ParcelaDTO{{Numero: 1, ValorAmortizacao: 180, ValorJuros: 16.11, ValorPrestacao: 196.11}, {Numero: 2, ValorAmortizacao: 180, ValorJuros: 12.89, ValorPrestacao: 192.89}, {Numero: 3, ValorAmortizacao: 180, ValorJuros: 9.67, ValorPrestacao: 189.67}, {Numero: 4, ValorAmortizacao: 180, ValorJuros: 6.44, ValorPrestacao: 186.44}, {Numero: 5, ValorAmortizacao: 180, ValorJuros: 3.22, ValorPrestacao: 183.22}}
	resp := &domain.ResultadoSimulacaoDTO{Tipo: "SAC", Parcelas: &items}

	require.Equal(t, value, resp)

}
