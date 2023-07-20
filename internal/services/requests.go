package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetExchangeRates(dateString string) ([]byte, error) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://www.cbr.ru/scripts/XML_daily.asp?date_req=%s",
		date.Format("02/01/2006"))
	fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
