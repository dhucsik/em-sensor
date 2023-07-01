# Decoder for EMS-300-MCS

Предположим, датчик прислал последовательность "0367F600046882060001". Нужно написать функцию, которая декодирует эту посылку и вернет значения вида:

Temperature: 24.6 C

Humidity: 65%

MagneticStatus: Open

## How to run tests

```
go test ./decoder
```