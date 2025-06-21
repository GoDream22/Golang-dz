// Конвертер валют

package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("___Конвертер валют___")
	fmt.Println("Доступные валюты: usd, eur, rub")

	currencyMap := map[string]float64{}
	var currency2 string
	var quantity float64
	var convertionResult float64

	for {
		var err error
		currencyMap, err = getFirstCurrency()
		if err != nil {
			fmt.Println(err)
		} else {
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

	availableSecondCurrency(currencyMap)

	for {
		var err error
		currency2, err = getSecondCurrency(currencyMap)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Валюта 2:", currency2)
			break
		}
	}

	convertionResult = currencyMap[currency2] * quantity

	fmt.Printf("Вы получите %.2f в валюте %s ", convertionResult, currency2)

}

func getFirstCurrency() (map[string]float64, error) {
	var input1 string
	usdConvert := map[string]float64{"rub": 80.77, "eur": 0.896}
	rubConvert := map[string]float64{"usd": 0.012, "eur": 0.011}
	eurConvert := map[string]float64{"usd": 1.12, "rub": 90.2}

	err := errors.New("Конвертируемая валюта введена с ошибкой")
	fmt.Print("Введите валюту, которую хотите конвертировать: ")
	fmt.Scan(&input1)
	input1 = strings.ToLower(input1)
	switch input1 {
	case "eur":
		return eurConvert, nil
	case "usd":
		return usdConvert, nil
	case "rub":
		return rubConvert, nil
	default:
		return nil, err
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

func availableSecondCurrency(currencyMap map[string]float64) {
	for currency, _ := range currencyMap {
		fmt.Println("Доступна валюта: ", currency)
	}
}

func getSecondCurrency(currencyMap map[string]float64) (string, error) {
	var input2 string
	err := errors.New("Получаемая валюта введена с ошибкой")
	fmt.Print("Введите валюту, которую хотите получить: ")
	fmt.Scan(&input2)
	input2 = strings.ToLower(input2)
	for currency, _ := range currencyMap {
		if input2 == currency {
			return input2, nil
		}
	}
	return "", err
}
