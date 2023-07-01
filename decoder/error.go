package decoder

import "errors"

var (
	ErrTemperatureNotFound    = errors.New("temperature channel, type not found")
	ErrHumidityNotFound       = errors.New("humidity channel, type not found")
	ErrMagneticStatusNotFound = errors.New("magnetic status channel, type not found")
)
