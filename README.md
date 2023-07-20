# Курс валют
Реализовать консольную утилиту, которая выводит курсы валют ЦБ РФ за определенную дату. Для получения курсов необходимо использовать официальный API ЦБ РФ https://www.cbr.ru/development/sxml/.

Для реализации задачи вы можете использовать любые сторонние библиотеки, которые считаете необходимыми.

## Интерфейс:

currency_rates --code=USD --date=2022-10-08
 
Описание параметров:
code - код валюты в формате ISO 4217
date - дата в формате YYYY-MM-DD
Вывод:

USD (Доллар США): 61,2475

## Запуск

```
go mod download
go build github.com/loveletter4you/exchange_rates/cmd/currency_rates
./currency_rates -code=$your_code -date=$your_date
```
