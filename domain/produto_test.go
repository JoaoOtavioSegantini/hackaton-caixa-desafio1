package domain_test

import (
	"testing"

	"github.com/hack-caixa/domain"
	"github.com/stretchr/testify/require"
)

func TestProductValidationWhenProductIsNotValid(t *testing.T) {
	produto := domain.NewProduct()

	err := produto.Validate()

	require.Error(t, err)

	produto.CO_PRODUTO = 1

	err = produto.Validate()

	require.Error(t, err)

	produto.PC_TAXA_JUROS = 0.0175

	err = produto.Validate()

	require.Error(t, err)

	produto.NO_PRODUTO = "Produto 1"

	err = produto.Validate()

	require.Error(t, err)

}

func TestProductValidationWhenProductIsValid(t *testing.T) {
	produto := domain.NewProduct()
	produto.CO_PRODUTO = 1
	produto.NO_PRODUTO = "Produto 1"
	produto.PC_TAXA_JUROS = 0.0175
	produto.VR_MINIMO = 200.00

	err := produto.Validate()

	require.Nil(t, err)
}
