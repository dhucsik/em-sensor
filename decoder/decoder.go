package decoder

import (
	"fmt"
	"strconv"
	"strings"
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
	tempIndex := strings.Index(hex, temperatureChannelType)
	if tempIndex < 0 {
		return nil, ErrTemperatureNotFound
	}

	tempValue, err := hexToDecimal(fmt.Sprintf("%s%s", hex[tempIndex+6:tempIndex+8], hex[tempIndex+4:tempIndex+6]))
	if err != nil {
		return nil, err
	}

	temperature := float64(tempValue) * 0.1

	humIndex := strings.Index(hex, humidityChannelType)
	if humIndex < 0 {
		return nil, ErrHumidityNotFound
	}

	humValue, err := hexToDecimal(hex[humIndex+4 : humIndex+6])
	if err != nil {
		return nil, err
	}

	humidity := float64(humValue) * 0.5

	magsIndex := strings.Index(hex, magneticStatusChannelType)
	if magsIndex < 0 {
		return nil, ErrMagneticStatusNotFound
	}

	magneticStatus := ""
	if hex[magsIndex+4:magsIndex+6] == "00" {
		magneticStatus = "Close"
	} else if hex[magsIndex+4:magsIndex+6] == "01" {
		magneticStatus = "Open"
	} else {
		return nil, ErrMagneticStatusNotFound
	}

	return &Result{
		Temperature:    temperature,
		Humidity:       humidity,
		MagneticStatus: magneticStatus,
	}, nil
}

func hexToDecimal(hex string) (int, error) {
	decimal, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return 0, err
	}

	return int(decimal), err
}
