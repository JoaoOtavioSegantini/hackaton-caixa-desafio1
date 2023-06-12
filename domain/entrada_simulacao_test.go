package domain_test

import (
	"testing"

	"github.com/hack-caixa/domain"
	"github.com/stretchr/testify/require"
)

func TestInputValidationWhenInputIsNotValid(t *testing.T) {
	input := domain.NewInput()

	err := input.Validate()

	require.Error(t, err)

	input.Prazo = 1

	err = input.Validate()

	require.Error(t, err)

	in := domain.NewInput()

	in.ValorDesejado = 500.00

	err = in.Validate()

	require.Error(t, err)

}

func TestInputValidationWhenInputIsValid(t *testing.T) {
	input := domain.NewInput()
	input.Prazo = 1
	input.ValorDesejado = 200.00

	err := input.Validate()

	require.Nil(t, err)
}
