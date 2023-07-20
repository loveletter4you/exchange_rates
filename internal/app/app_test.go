package app_test

import (
	"github.com/loveletter4you/exchange_rates/internal/app"
	"github.com/loveletter4you/exchange_rates/internal/model"
	"testing"
)

func TestAdd(t *testing.T) {
	got := app.App("USD", "2002-03-02")[0]
	want := model.Valute{
		NumCode:  840,
		CharCode: "USD",
		Nominal:  1,
		Name:     "Доллар США",
		Value:    "30,9436",
	}
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
