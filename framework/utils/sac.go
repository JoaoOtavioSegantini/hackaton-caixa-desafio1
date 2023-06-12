package utils

import (
	"fmt"
	"strconv"
)

func CalcularAmortizacaoSAC(valorEmprestimo float32, numParcelas int) float32 {
	valorParcela := valorEmprestimo / float32(numParcelas)
	return valorParcela
}

func CalcularParcelaSAC(juros float32, amortizacao float32) float32 {
	return amortizacao + juros
}

func CalcularJurosSAC(loanAmount, interestRate float64, loanTermMonths int) []float64 {
	monthlyRate := interestRate
	monthlyPayment := loanAmount / float64(loanTermMonths)
	var totalInterest []float64

	for i := 0; i < loanTermMonths; i++ {
		interest := loanAmount * monthlyRate
		pres := fmt.Sprintf("%.2f", interest)
		juros, _ := strconv.ParseFloat(pres, 64)
		totalInterest = append(totalInterest, juros)
		loanAmount -= monthlyPayment
	}

	return totalInterest
}
