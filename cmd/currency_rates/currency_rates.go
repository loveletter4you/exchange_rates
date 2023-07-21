package main

import (
	"flag"
	"fmt"
	"github.com/loveletter4you/exchange_rates/internal/app"
	"os"
	"time"
)

func main() {
	code := flag.String("code", "", "Currency code")
	date := flag.String("date", time.Now().Format("2006-01-02"), "Exchange date")
	flag.Parse()
	result := app.App(*code, *date)
	for _, valute := range result {
		_, err := fmt.Fprintf(os.Stdout, "%s (%s): %s\n",
			valute.CharCode, valute.Name, valute.Value)
		if err != nil {
			panic(err)
		}
	}
}
