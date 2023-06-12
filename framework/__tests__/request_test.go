package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hack-caixa/application/repositories"
	"github.com/hack-caixa/domain"
	"github.com/hack-caixa/framework/config/database"
	"github.com/hack-caixa/framework/routes"
	"github.com/stretchr/testify/assert"
)

func TestMakeSimulation(t *testing.T) {

	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)
	db := database.NewDbTest()

	repositories.SetupDatabase(db)

	routes.SetupRouter(r, db)

	simulacao := domain.EntradaSimulacaoDTO{ValorDesejado: 900.00, Prazo: 5}
	jsonValue, _ := json.Marshal(simulacao)

	req, err := http.NewRequestWithContext(ctx, "POST", "/api/Simulacao/", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Errorf("got error: %s", err)
	}
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())
	assert.Contains(t, w.Body.String(), "{\"codigoProduto\":1,\"descricaoProduto\":\"Produto 1\",\"taxaJuros\":0.0179,\"ResultadoSimulacao\":[{\"tipo\":\"SAC\",\"parcelas\":[{\"numero\":1,\"valorAmortizacao\":180,\"valorJuros\":16.11,\"valorPrestacao\":196.11},{\"numero\":2,\"valorAmortizacao\":180,\"valorJuros\":12.89,\"valorPrestacao\":192.89},{\"numero\":3,\"valorAmortizacao\":180,\"valorJuros\":9.67,\"valorPrestacao\":189.67},{\"numero\":4,\"valorAmortizacao\":180,\"valorJuros\":6.44,\"valorPrestacao\":186.44},{\"numero\":5,\"valorAmortizacao\":180,\"valorJuros\":3.22,\"valorPrestacao\":183.22}]},{\"tipo\":\"PRICE\",\"parcelas\":[{\"numero\":1,\"valorAmortizacao\":173.67,\"valorJuros\":16.11,\"valorPrestacao\":189.78},{\"numero\":2,\"valorAmortizacao\":176.78,\"valorJuros\":13,\"valorPrestacao\":189.78},{\"numero\":3,\"valorAmortizacao\":179.94,\"valorJuros\":9.84,\"valorPrestacao\":189.78},{\"numero\":4,\"valorAmortizacao\":183.16,\"valorJuros\":6.62,\"valorPrestacao\":189.78},{\"numero\":5,\"valorAmortizacao\":186.44,\"valorJuros\":3.34,\"valorPrestacao\":189.78}]}]}")

}
