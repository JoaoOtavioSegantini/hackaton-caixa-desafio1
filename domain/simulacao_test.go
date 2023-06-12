package domain_test

import (
	"testing"

	"github.com/hack-caixa/domain"
	"github.com/stretchr/testify/require"
)

func TestSimulacaoValidation(t *testing.T) {
	simulacao := domain.NewSimulacao()

	err := simulacao.Validate()

	require.Error(t, err)

	simulacao.CodigoProduto = 1

	err = simulacao.Validate()

	require.Error(t, err)

	simulacao.DescricaoProduto = "simple description"

	err = simulacao.Validate()

	require.Error(t, err)

	simulacao.TaxaJuros = 54.98

	err = simulacao.Validate()

	require.Nil(t, err)

}
