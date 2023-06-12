package domain_test

import (
	"testing"

	"github.com/hack-caixa/domain"
	"github.com/stretchr/testify/require"
)

func TestParcelaValidation(t *testing.T) {
	parcela := domain.NewParcelaDTO()

	err := parcela.Validate()

	require.Error(t, err)

	parcela.Numero = 1

	err = parcela.Validate()

	require.Error(t, err)

	parcela.ValorAmortizacao = 500.98

	err = parcela.Validate()

	require.Error(t, err)

	parcela.ValorJuros = 54.98

	err = parcela.Validate()

	require.Error(t, err)

	parcela.ValorPrestacao = 54.98

	err = parcela.Validate()

	require.Nil(t, err)

}
