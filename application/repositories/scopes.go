package repositories

import "gorm.io/gorm"

func MaxMeses(Prazo int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("nu_maximo_meses >= ? OR nu_maximo_meses IS NULL", Prazo)
	}
}

func MinMeses(Prazo int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("nu_minimo_meses <= ?", Prazo)
	}
}

func ValorMinimo(ValorDesejado float32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("vr_minimo <= ?", ValorDesejado)
	}
}

func ValorMaximo(ValorDesejado float32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("vr_maximo >= ? OR vr_maximo IS NULL", ValorDesejado)
	}
}
