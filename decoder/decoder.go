package decoder

import (
	"fmt"
	"strconv"
)

type Result struct {
	Temperature    float64
	Humidity       float64
	MagneticStatus string
}

const (
	temperatureChannelType    = "0367"
	humidityChannelType       = "0468"
	magneticStatusChannelType = "0600"
)

func Decode(hex string) (*Result, error) {
	if len(hex) != 20 {
		return nil, ErrLenNotCorrect
	}

	var (
		temperature    float64
		humidity       float64
		magneticStatus string
		err            error
	)

Loop:
	for {
		if len(hex) < 4 {
			return nil, ErrDataFormat
		}

		switch hex[0:4] {
		case temperatureChannelType:
			if len(hex) < 8 {
				return nil, ErrDataFormat
			}

			temperature, err = hexToTemperature(hex[0:8])
			if err != nil {
				return nil, err
			}

			if len(hex) == 8 {
				break Loop
			}

			hex = hex[8:]

		case humidityChannelType:
			if len(hex) < 6 {
				return nil, ErrDataFormat
			}

			humidity, err = hexToHumidity(hex[0:6])
			if err != nil {
				return nil, err
			}

			if len(hex) == 6 {
				break Loop
			}

			hex = hex[6:]

		case magneticStatusChannelType:
			if len(hex) < 6 {
				return nil, ErrDataFormat
			}

			magneticStatus, err = hexToMagneticStatus(hex[0:6])
			if err != nil {
				return nil, err
			}

			if len(hex) == 6 {
				break Loop
			}

			hex = hex[6:]

		default:
			return nil, ErrDataFormat
		}
	}

	return &Result{
		Temperature:    temperature,
		Humidity:       humidity,
		MagneticStatus: magneticStatus,
	}, nil
}

func hexToTemperature(hex string) (float64, error) {
	tempValue, err := hexToDecimal(fmt.Sprintf("%s%s", hex[6:8], hex[4:6]))
	if err != nil {
		return 0, err
	}

	temperature := float64(tempValue) * 0.1

	return temperature, nil
}

func hexToHumidity(hex string) (float64, error) {
	humValue, err := hexToDecimal(hex[4:6])
	if err != nil {
		return 0, err
	}

	humidity := float64(humValue) * 0.5
	return humidity, nil
}

func hexToMagneticStatus(hex string) (string, error) {
	magneticStatus := ""
	if hex[4:6] == "00" {
		magneticStatus = "Close"
	} else if hex[4:6] == "01" {
		magneticStatus = "Open"
	} else {
		return magneticStatus, ErrDataFormat
	}

	return magneticStatus, nil
}

func hexToDecimal(hex string) (int, error) {
	decimal, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return 0, err
	}

	if len(hex) == 4 && hex[0] >= '8' {
		decimal -= 65536
	}

	return int(decimal), err
}
