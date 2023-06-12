package utils

import (
	"github.com/hack-caixa/domain"
)

func MakeSacParcelasDTO(sac []domain.ParcelaDTO, in domain.EntradaSimulacaoDTO, produto domain.PRODUTO) *domain.ResultadoSimulacaoDTO {
	for i := range sac {

		amortizacaoSac := CalcularAmortizacaoSAC(in.ValorDesejado, int(in.Prazo))
		jurosSac := CalcularJurosSAC(float64(in.ValorDesejado), float64(produto.PC_TAXA_JUROS), int(in.Prazo))
		sac[i] = domain.ParcelaDTO{
			Numero:           int32(i + 1),
			ValorAmortizacao: amortizacaoSac,
			ValorJuros:       float32(jurosSac[i]),
			ValorPrestacao:   CalcularParcelaSAC(float32(jurosSac[i]), amortizacaoSac),
		}

	}

	resultadoSac := &domain.ResultadoSimulacaoDTO{Tipo: "SAC", Parcelas: &sac}

	return resultadoSac
}
