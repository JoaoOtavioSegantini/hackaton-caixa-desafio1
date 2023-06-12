package repositories_test

import (
	"context"
	"testing"

	"github.com/hack-caixa/application/repositories"
	"github.com/hack-caixa/domain"
	"github.com/hack-caixa/framework/config/database"
	"github.com/stretchr/testify/require"
)

func TestProductRepository(t *testing.T) {
	db := database.NewDbTest()

	repo := repositories.NewProductRepository(db)

	repositories.SetupDatabase(db)

	input := domain.NewInput()

	input.Prazo = 5
	input.ValorDesejado = 900.00
	produto_repo, err := repo.FindProductScoped(context.TODO(), input)

	require.Nil(t, err)

	require.Equal(t, produto_repo.CO_PRODUTO, 1)
	require.Equal(t, produto_repo.NO_PRODUTO, "Produto 1")
	require.Equal(t, produto_repo.NU_MINIMO_MESES, 0)
	require.Equal(t, produto_repo.NU_MAXIMO_MESES, 24)
	require.Equal(t, produto_repo.PC_TAXA_JUROS, float32(0.017900000))
	require.Equal(t, produto_repo.VR_MINIMO, float32(200.00))
	require.Equal(t, produto_repo.VR_MAXIMO, float32(10000.00))

	input.Prazo = 25
	input.ValorDesejado = 15000.00

	produto_repo, err = repo.FindProductScoped(context.TODO(), input)

	require.Nil(t, err)

	require.Equal(t, produto_repo.CO_PRODUTO, 2)
	require.Equal(t, produto_repo.NO_PRODUTO, "Produto 2")
	require.Equal(t, produto_repo.NU_MINIMO_MESES, 25)
	require.Equal(t, produto_repo.NU_MAXIMO_MESES, 48)
	require.Equal(t, produto_repo.PC_TAXA_JUROS, float32(0.017500000))
	require.Equal(t, produto_repo.VR_MINIMO, float32(10001.00))
	require.Equal(t, produto_repo.VR_MAXIMO, float32(100000.00))

	input.Prazo = 49
	input.ValorDesejado = 115000.01

	produto_repo, err = repo.FindProductScoped(context.TODO(), input)

	require.Nil(t, err)

	require.Equal(t, produto_repo.CO_PRODUTO, 3)
	require.Equal(t, produto_repo.NO_PRODUTO, "Produto 3")
	require.Equal(t, produto_repo.NU_MINIMO_MESES, 49)
	require.Equal(t, produto_repo.NU_MAXIMO_MESES, 96)
	require.Equal(t, produto_repo.PC_TAXA_JUROS, float32(0.018200000))
	require.Equal(t, produto_repo.VR_MINIMO, float32(100000.01))
	require.Equal(t, produto_repo.VR_MAXIMO, float32(1000000.00))

	input.Prazo = 100
	input.ValorDesejado = 1005000.80

	produto_repo, err = repo.FindProductScoped(context.TODO(), input)

	require.Nil(t, err)
	require.Equal(t, produto_repo.CO_PRODUTO, 4)
	require.Equal(t, produto_repo.NO_PRODUTO, "Produto 4")
	require.Equal(t, produto_repo.NU_MINIMO_MESES, 96)
	require.Equal(t, produto_repo.NU_MAXIMO_MESES, int(0))
	require.Equal(t, produto_repo.PC_TAXA_JUROS, float32(0.015100000))
	require.Equal(t, produto_repo.VR_MINIMO, float32(1000000.01))
	require.Equal(t, produto_repo.VR_MAXIMO, float32(0))

}
