package decoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		name   string
		hex    string
		expect *Result
		expErr bool
		err    error
	}{
		{
			name: "success1",
			hex:  "0367F600046882060001",
			expect: &Result{
				Temperature:    24.6,
				Humidity:       65.0,
				MagneticStatus: "Open"},
			expErr: false,
			err:    nil,
		},
		{
			name: "success2",
			hex:  "0468600600000367AF00",
			expect: &Result{
				Temperature:    17.5,
				Humidity:       48.0,
				MagneticStatus: "Close",
			},
			expErr: false,
			err:    nil,
		},
		{
			name:   "no temperature",
			hex:    "046882060001",
			expect: nil,
			expErr: true,
			err:    ErrTemperatureNotFound,
		},
		{
			name:   "no humidity",
			hex:    "0600000367AF00",
			expect: nil,
			expErr: true,
			err:    ErrHumidityNotFound,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			result, err := Decode(tCase.hex)
			assert.Equal(t, result, tCase.expect, "they should be equal")
			if tCase.expErr {
				assert.EqualError(t, tCase.err, err.Error(), "error must be nil")
			} else {
				assert.NoError(t, err, "there must be no errors")
			}
		})
	}

}
