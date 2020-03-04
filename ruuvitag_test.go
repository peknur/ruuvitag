package ruuvitag

import (
	"encoding/hex"
	"testing"
)

func TestIsRuuvitag(t *testing.T) {
	data, err := hex.DecodeString("99040300FF6300008001800180010000")
	if err != nil {
		t.Fatal(err)
	}
	if isRuuviDevice(data) == false {
		t.Errorf("isRuuviDevice failed")
	}
	data, err = hex.DecodeString("99010300FF6300008001800180010000")
	if err != nil {
		t.Fatal(err)
	}
	if isRuuviDevice(data) == true {
		t.Errorf("isRuuviDevice failed with invalid data")
	}
	data = make([]byte, 16)
	if isRuuviDevice(data) == true {
		t.Errorf("isRuuviDevice failed with empty []byte array")
	}
	data = make([]byte, 1)
	if isRuuviDevice(data) == true {
		t.Errorf("isRuuviDevice failed with empty []byte array")
	}
}
