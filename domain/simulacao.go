package domain

import "github.com/asaskevich/govalidator"

type SimulacaoDTO struct {
	CodigoProduto      int                      `json:"codigoProduto" valid:"notnull"`
	DescricaoProduto   string                   `json:"descricaoProduto" valid:"-"`
	TaxaJuros          float32                  `json:"taxaJuros" valid:"notnull"`
	ResultadoSimulacao *[]ResultadoSimulacaoDTO `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (in *SimulacaoDTO) Validate() error {
	_, err := govalidator.ValidateStruct(in)

	if err != nil {
		return err
	}
	return nil
}

func NewSimulacao() *SimulacaoDTO {

	return &SimulacaoDTO{}
}
