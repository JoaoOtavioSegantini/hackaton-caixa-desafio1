package domain

import "github.com/asaskevich/govalidator"

type ParcelaDTO struct {
	Numero           int32   `json:"numero" valid:"notnull"`
	ValorAmortizacao float32 `json:"valorAmortizacao" valid:"notnull"`
	ValorJuros       float32 `json:"valorJuros" valid:"notnull"`
	ValorPrestacao   float32 `json:"valorPrestacao" valid:"notnull"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (in *ParcelaDTO) Validate() error {
	_, err := govalidator.ValidateStruct(in)

	if err != nil {
		return err
	}
	return nil
}

func NewParcelaDTO() *ParcelaDTO {
	return &ParcelaDTO{}
}
