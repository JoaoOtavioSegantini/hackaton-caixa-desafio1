package domain

import "github.com/asaskevich/govalidator"

type EntradaSimulacaoDTO struct {
	ValorDesejado float32 `json:"valorDesejado" valid:"notnull"`
	Prazo         int32   `json:"prazo" valid:"notnull"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (in *EntradaSimulacaoDTO) Validate() error {
	_, err := govalidator.ValidateStruct(in)

	if err != nil {
		return err
	}
	return nil
}

func NewInput() *EntradaSimulacaoDTO {
	return &EntradaSimulacaoDTO{}
}
