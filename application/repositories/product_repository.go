package repositories

import (
	"context"

	"github.com/hack-caixa/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProductScoped(ctx context.Context, input *domain.EntradaSimulacaoDTO) (*domain.PRODUTO, error)
}

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryDb {
	return &ProductRepositoryDb{Db: db}
}

func (repo ProductRepositoryDb) FindProductScoped(ctx context.Context, input *domain.EntradaSimulacaoDTO) (*domain.PRODUTO, error) {

	var produto domain.PRODUTO

	err := repo.Db.Scopes(
		ValorMinimo(input.ValorDesejado), MinMeses(input.Prazo),
		ValorMaximo(input.ValorDesejado), MaxMeses(input.Prazo)).Find(&produto).Error

	if err != nil {
		return nil, err
	}

	err = produto.Validate()

	if err != nil {
		return nil, err
	}

	return &produto, nil

}

func SetupDatabase(db *gorm.DB) {
	produto := domain.NewProduct()
	produto.CO_PRODUTO = 1
	produto.NO_PRODUTO = "Produto 1"
	produto.NU_MAXIMO_MESES = 24
	produto.NU_MINIMO_MESES = 0
	produto.PC_TAXA_JUROS = 0.017900000
	produto.VR_MINIMO = 200.00
	produto.VR_MAXIMO = 10000.00

	db.Create(produto)

	produto2 := domain.NewProduct()
	produto2.CO_PRODUTO = 2
	produto2.NO_PRODUTO = "Produto 2"
	produto2.NU_MAXIMO_MESES = 48
	produto2.NU_MINIMO_MESES = 25
	produto2.PC_TAXA_JUROS = 0.017500000
	produto2.VR_MINIMO = 10001.00
	produto2.VR_MAXIMO = 100000.00

	db.Create(produto2)

	produto3 := domain.NewProduct()
	produto3.CO_PRODUTO = 3
	produto3.NO_PRODUTO = "Produto 3"
	produto3.NU_MAXIMO_MESES = 96
	produto3.NU_MINIMO_MESES = 49
	produto3.PC_TAXA_JUROS = 0.018200000
	produto3.VR_MINIMO = 100000.01
	produto3.VR_MAXIMO = 1000000.00

	db.Create(produto3)

	produto4 := domain.NewProduct()
	produto4.CO_PRODUTO = 4
	produto4.NO_PRODUTO = "Produto 4"
	produto4.NU_MINIMO_MESES = 96
	produto4.PC_TAXA_JUROS = 0.015100000
	produto4.VR_MINIMO = 1000000.01

	db.Create(produto4)

}
