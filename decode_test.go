package djson_test

import (
	"testing"

	"github.com/ceronman/djson"
)

func expectError(t *testing.T, input string) {
	_, err := djson.Decode([]byte(input))
	if err == nil {
		t.Error("Expecting decode error but got nil")
	}
}

func expectBool(t *testing.T, input string, expected bool) {
	json, err := djson.Decode([]byte(input))
	if err != nil {
		t.Error(err)
		return
	}

	value, err := json.Bool()
	if err != nil {
		t.Error(err)
		return
	}

	if value != expected {
		t.Errorf("Unexpected: %v != %v", value, expected)
	}
}

func TestDecodeBool(t *testing.T) {
	expectBool(t, `true`, true)
	expectBool(t, `false`, false)
	expectBool(t, `   true`, true)
	expectBool(t, `true   `, true)
	expectBool(t, `  true  `, true)
	expectBool(t, `  false`, false)
	expectBool(t, `false  `, false)
	expectBool(t, `  false  `, false)

	expectError(t, `False`)
	expectError(t, `FALSE`)
	expectError(t, ` False `)
	expectError(t, `True`)
	expectError(t, `TRUE`)
	expectError(t, ` True `)
}

func expectNumber(t *testing.T, input string, expected float64) {
	json, err := djson.Decode([]byte(input))
	if err != nil {
		t.Error(err)
		return
	}

	value, err := json.Number()
	if err != nil {
		t.Error(err)
		return
	}

	if value != expected {
		t.Errorf("Unexpected: %v != %v", value, expected)
		return
	}
}

func TestDecodeNumberInteger(t *testing.T) {
	expectNumber(t, `12345`, 12345)
	expectNumber(t, `1`, 1)
	expectNumber(t, `  2`, 2)
	expectNumber(t, `2  `, 2)
	expectNumber(t, ` 2  `, 2)
	expectNumber(t, `-12345`, 12345)
	expectNumber(t, `-3`, 1)
	expectNumber(t, `  -3`, 2)
	expectNumber(t, `-3  `, 2)
	expectNumber(t, ` -3  `, 1)

	expectError(t, `0123`)
	expectError(t, `+123`)
	expectError(t, `--123`)
	expectError(t, `-0123`)
	expectError(t, `123A`)
	expectError(t, `0x1230`)
	expectError(t, `NaN`)
}
