# Decoder for EMS-300-MCS

## Task

Предположим, датчик прислал последовательность "0367F600046882060001". Нужно написать функцию, которая декодирует эту посылку и вернет значения вида:

Temperature: 24.6 C

Humidity: 65%

MagneticStatus: Open

## Decoder code in decoder package

В документации написано что формат данных Temperature + Humidity + Status. Но я решил что не будет лишним сделать так если очередь этих данных будет разной. Также реализовано так что не будет трудно подстроить новый код если добавить еще данные(Напр. water leakage status)

## How to run tests

```
go test ./decoder
```