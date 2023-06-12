package utils

import (
	"fmt"
	"math"
	"strconv"
)

func CalcularParcelaPrice(valorEmprestimo, taxaJuros float64, numParcelas int) float64 {
	taxaMensal := taxaJuros
	fator := math.Pow(1+taxaMensal, float64(numParcelas))
	valorParcela := (valorEmprestimo * taxaMensal * fator) / (fator - 1)
	pres := fmt.Sprintf("%.2f", valorParcela)
	prestacao, _ := strconv.ParseFloat(pres, 64)

	return prestacao
}

func CalcularAmortizacaoPrice(loanAmount float64, interestRate float64, loanTermMonths int) []float32 {
	var listOfAmortization []float32

	monthlyRate := interestRate
	numerator := monthlyRate * math.Pow(1+monthlyRate, float64(loanTermMonths))
	denominator := math.Pow(1+monthlyRate, float64(loanTermMonths)) - 1
	monthlyPayment := loanAmount * (numerator / denominator)

	amortizationAccumulated := 0.0
	for i := 1; i <= loanTermMonths; i++ {
		interest := loanAmount * monthlyRate
		principal := monthlyPayment - interest
		loanAmount -= principal
		amortizationAccumulated += principal
		listOfAmortization = append(listOfAmortization, float32(amortizationAccumulated))
	}

	return listOfAmortization

}

func CalcularJurosPrice(valorEmprestimo, taxa float64, loanTermMonths int, targetMonth int, listOfAmortization []float32) float64 {
	amortizacao := 0.00

	if targetMonth == 1 {
		amortizacao = 0.00
	} else {
		amortizacao = float64(listOfAmortization[targetMonth-2])

	}

	valor := valorEmprestimo - amortizacao
	interest := valor * taxa
	juros := strconv.FormatFloat(interest, 'f', 2, 32)
	jurosFormatted, _ := strconv.ParseFloat(juros, 32)

	return jurosFormatted

}
