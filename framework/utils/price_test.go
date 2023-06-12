package utils_test

import (
	"testing"

	"github.com/hack-caixa/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestCalcularParcelaPriceFunction(t *testing.T) {
	valor := 900.00
	juros := 0.017900000
	num := 5
	value := utils.CalcularParcelaPrice(valor, juros, num)

	require.Equal(t, value, 189.78)
}

func TestCalcularAmortizacaoPriceFunction(t *testing.T) {
	valor := 900.00
	juros := 0.017900000
	num := 5
	value := utils.CalcularAmortizacaoPrice(valor, juros, num)

	require.Equal(t, value, []float32{173.67032, 350.4493, 530.3927, 713.557, 900})

}

func TestCalcularJurosPriceFunction(t *testing.T) {
	valor := 900.00
	juros := 0.017900000
	num := 5
	value := utils.CalcularJurosPrice(valor, juros, num, 1, []float32{173.67032, 350.4493, 530.3927, 713.557, 900})

	require.Equal(t, value, float64(16.110000610351562))

	value = utils.CalcularJurosPrice(valor, juros, num, 2, []float32{173.67032, 350.4493, 530.3927, 713.557, 900})

	require.Equal(t, value, float64(13))

	value = utils.CalcularJurosPrice(valor, juros, num, 3, []float32{173.67032, 350.4493, 530.3927, 713.557, 900})

	require.Equal(t, value, float64(9.84000015258789))

}
