package app

import (
	"github.com/loveletter4you/exchange_rates/internal/model"
	"github.com/loveletter4you/exchange_rates/internal/services"
)

func App(code string, date string) []model.Valute {
	resp, err := services.GetExchangeRates(date)
	if err != nil {
		panic(err)
	}
	valutes, err := services.ParseDaily(resp)
	if err != nil {
		panic(err)
	}
	var result []model.Valute
	for _, valute := range valutes.Valute {
		if (code == "") || (code == valute.CharCode) {
			result = append(result, valute)
			if err != nil {
				panic(err)
			}
		}
	}
	return result
}
