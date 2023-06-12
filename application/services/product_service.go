package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hack-caixa/application/repositories"
	"github.com/hack-caixa/domain"
	"github.com/hack-caixa/framework/queue"
	"github.com/hack-caixa/framework/utils"
	"gorm.io/gorm"
)

func MakeSimulation(ctx *gin.Context, db *gorm.DB) {
	var (
		items []domain.ResultadoSimulacaoDTO
		input domain.EntradaSimulacaoDTO
	)

	err := ctx.BindJSON(&input)

	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return

	}

	err = input.Validate()

	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return

	}

	conn := repositories.NewProductRepository(db)
	produto, err := conn.FindProductScoped(ctx, &input)

	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	sac := make([]domain.ParcelaDTO, input.Prazo)
	price := make([]domain.ParcelaDTO, input.Prazo)

	resultadoSac := utils.MakeSacParcelasDTO(sac, input, *produto)
	resultadoPrice := utils.MakePriceParcelasDTO(price, input, *produto)
	items = append(items, *resultadoSac, *resultadoPrice)

	response := &domain.SimulacaoDTO{
		CodigoProduto:      produto.CO_PRODUTO,
		DescricaoProduto:   produto.NO_PRODUTO,
		TaxaJuros:          produto.PC_TAXA_JUROS,
		ResultadoSimulacao: &items,
	}

	go queue.PublishInEventHub(ctx, *response, db)

	ctx.JSON(http.StatusOK, &response)
}
