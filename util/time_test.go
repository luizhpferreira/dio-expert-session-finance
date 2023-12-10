package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2019-02-12T10:08:02")

	if convertedTime.Year() != 2019 {
		t.Errorf("espera que o ano seja 2019")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("espera que o ano seja 02")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("espera que o ano seja 10")
	}
}
