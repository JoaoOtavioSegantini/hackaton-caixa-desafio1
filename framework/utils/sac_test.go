package utils_test

import (
	"testing"

	"github.com/hack-caixa/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestCalcularAmortizacaoSACFunction(t *testing.T) {
	valor := 900.00
	num := 5
	value := utils.CalcularAmortizacaoSAC(float32(valor), num)

	require.Equal(t, value, float32(180))

}

func TestCalcularParcelaSACFunction(t *testing.T) {
	valor := 900.00
	amortizacao := 235.00
	value := utils.CalcularParcelaSAC(float32(valor), float32(amortizacao))

	require.Equal(t, value, float32(1135))

}

func TestCalcularJurosSACFunction(t *testing.T) {
	valor := 900.00
	amortizacao := 235.67
	num := 5

	value := utils.CalcularJurosSAC(valor, amortizacao, num)

	require.Equal(t, value, []float64{212103, 169682.4, 127261.8, 84841.2, 42420.6})

}
