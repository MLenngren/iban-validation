package ibanValidator_test

import (
	"testing"

	validator "github.com/mlenngren/iban-validator/ibanValidator"
	"github.com/stretchr/testify/assert"
)

func TestValidateIban(t *testing.T) {
	validatedResult := validator.ValidateIban("BE71096123456769")
	assert.Equal(t, 1, validatedResult)
}

func TestRotateFirst4Chars(t *testing.T) {
	rotatedResult := validator.RotateFirst4Chars("BE71096123456769")
	assert.Equal(t, "096123456769BE71", rotatedResult)
}

func TestCharacterToNumbers(t *testing.T) {
	result := validator.CharacterToNumbers("096123456769BE71")
	assert.Equal(t, "096123456769111471", result)
}

func TestLargeMod97Calc(t *testing.T) {
	result := validator.LargeMod97Calc("096123456769111471")
	assert.Equal(t, 1, result)
}
