package domain

type ResultadoSimulacaoDTO struct {
	Tipo     string        `json:"tipo" valid:"-"`
	Parcelas *[]ParcelaDTO `json:"parcelas" valid:"-"`
}
