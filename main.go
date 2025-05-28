// Конвертер валют

package main

import (
	"errors"
	"fmt"
	"strings"
)

var currency1 string
var currency2 string
var quantity float64

const usdToEur = 0.8955
const usdToRub = 80.77
const eurTorub = usdToRub / usdToEur

func main() {
	fmt.Println("___Конвертер валют___")
	fmt.Println("Доступные валюты: usd, eur, rub")

	for {
		var err error
		currency1, err = getFirstCurrency()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Валюта 1:", currency1)
			break
		}
	}
	for {
		var err error
		quantity, err = getQuantutyFirstCurrency()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Количество конвертируемой валюты:", quantity)
			break
		}
	}

	fmt.Println(availableSecondCurrency(currency1))

	for {
		var err error
		currency2, err = getSecondCurrency()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Валюта 2:", currency2)
			break
		}
	}

	fmt.Printf("Вы получите %.2f в валюте %s ", converter(), currency2)

}

func getFirstCurrency() (string, error) {
	var input1 string
	err := errors.New("Конвертируемая валюта введена с ошибкой")
	fmt.Print("Введите валюту, которую хотите конвертировать: ")
	fmt.Scan(&input1)
	switch strings.ToLower(input1) {
	case "eur":
		return input1, nil
	case "usd":
		return input1, nil
	case "rub":
		return input1, nil
	default:
		return "", err
	}
}

func getQuantutyFirstCurrency() (float64, error) {
	var inputQuantity float64
	err := errors.New("Количество валюты введено с ошибкой")
	fmt.Print("Введите количество конвертируемой валюты: ")
	fmt.Scan(&inputQuantity)
	if inputQuantity <= 0 {
		return 0, err
	}
	return inputQuantity, nil
}

func availableSecondCurrency(currency string) string {
	switch strings.ToLower(currency) {
	case "eur":
		return "Выберите usd или rub"
	case "usd":
		return "Выберите eur или rub"
	case "rub":
		return "Выберите usd или eur"
	default:
		return ""
	}
}

func getSecondCurrency() (string, error) {
	var input2 string
	err := errors.New("Получаемая валюта введена с ошибкой")
	fmt.Print("Введите валюту, которую хотите получить: ")
	fmt.Scan(&input2)
	if currency1 == "eur" {
		switch strings.ToLower(input2) {
		case "usd":
			return input2, nil
		case "rub":
			return input2, nil
		default:
			return "", err
		}
	} else if currency1 == "usd" {
		switch strings.ToLower(input2) {
		case "eur":
			return input2, nil
		case "rub":
			return input2, nil
		default:
			return "", err
		}
	} else {
		switch strings.ToLower(input2) {
		case "usd":
			return input2, nil
		case "eur":
			return input2, nil
		default:
			return "", err
		}
	}
}

func converter() float64 {
	var convertionResult float64
	if currency1 == "usd" && currency2 == "rub" {
		convertionResult = quantity * usdToRub
		return convertionResult
	} else if currency1 == "usd" && currency2 == "eur" {
		convertionResult = quantity * usdToEur
		return convertionResult
	} else if currency1 == "rub" && currency2 == "usd" {
		convertionResult = quantity / usdToRub
		return convertionResult
	} else if currency1 == "rub" && currency2 == "eur" {
		convertionResult = quantity / usdToRub * usdToEur
		return convertionResult
	} else if currency1 == "eur" && currency2 == "usd" {
		convertionResult = quantity / usdToEur
		return convertionResult
	} else {
		convertionResult = quantity * eurTorub
		return convertionResult
	}

}
