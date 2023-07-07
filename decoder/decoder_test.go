package decoder

import (
	"errors"
	"testing"
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
			name: "incorrect format",
			hex:  "0468600600000367AF00",
			expect: &Result{
				Temperature:    17.5,
				Humidity:       48,
				MagneticStatus: "Close",
			},
			expErr: false,
			err:    nil,
		},
		{
			name: "success2",
			hex:  "03670468046871060001",
			expect: &Result{
				Temperature:    2662.8,
				Humidity:       56.5,
				MagneticStatus: "Open",
			},
			expErr: false,
			err:    nil,
		},
		{
			name: "success3",
			hex:  "036738FF046842060000",
			expect: &Result{
				Temperature:    -20,
				Humidity:       33,
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
			err:    ErrLenNotCorrect,
		},
		{
			name:   "no humidity",
			hex:    "0600000367AF00",
			expect: nil,
			expErr: true,
			err:    ErrLenNotCorrect,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			result, err := Decode(tCase.hex)
			if tCase.expErr {
				if err == nil {
					t.Errorf("\"Decode('%s')\" FAILED, expected error -> %v, got nil error", tCase.hex, tCase.err)
				}
				if !errors.Is(err, tCase.err) {
					t.Errorf("\"Decode('%s')\" FAILED, expected error -> %v, got -> %v", tCase.hex, tCase.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("\"Decode('%s')\" FAILED, expected nil error, got -> %v", tCase.hex, err)
				}
				if *result != *tCase.expect {
					t.Errorf("\"Decode('%s')\" FAILED, expected output -> %v, got -> %v", tCase.hex, tCase.expect, result)
				}
			}

			t.Log("Succeded")
		})
	}

}
