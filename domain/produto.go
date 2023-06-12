package domain

import "github.com/asaskevich/govalidator"

type PRODUTO struct {
	CO_PRODUTO      int     `json:"CO_PRODUTO" valid:"notnull" gorm:"primary_key"`
	NO_PRODUTO      string  `json:"NO_PRODUTO" valid:"notnull"`
	PC_TAXA_JUROS   float32 `json:"PC_TAXA_JUROS" valid:"notnull"`
	NU_MINIMO_MESES int     `json:"NU_MINIMO_MESES" valid:"-"`
	NU_MAXIMO_MESES int     `json:"NU_MAXIMO_MESES" valid:"-" gorm:"default:null"`
	VR_MINIMO       float32 `json:"VR_MINIMO" valid:"notnull"`
	VR_MAXIMO       float32 `json:"VR_MAXIMO" valid:"-" gorm:"default:null"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (in *PRODUTO) Validate() error {
	_, err := govalidator.ValidateStruct(in)

	if err != nil {
		return err
	}
	return nil
}

func NewProduct() *PRODUTO {

	return &PRODUTO{}
}
