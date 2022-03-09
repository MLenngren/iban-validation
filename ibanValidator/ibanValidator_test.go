package ibanValidator_test

import (
	"testing"

	validator "github.com/mlenngren/iban-validator/ibanValidator"
	"github.com/stretchr/testify/assert"
)

func TestValidateIban(t *testing.T) {
	validatedResult := validator.ValidateIban("BE71096123456769")
	assert.Equal(t, 1, validatedResult)

	validatedResult = validator.ValidateIban("GB82WEST12345698765432")
	assert.Equal(t, 1, validatedResult)

	validatedResult = validator.ValidateIban("AL35202111090000000001234567")
	assert.Equal(t, 1, validatedResult)

	validatedResult = validator.ValidateIban("MU43BOMM0101123456789101000MUR")
	assert.Equal(t, 1, validatedResult)

	validatedResult = validator.ValidateIban("NI92BAMC000000000000000003123123")
	assert.Equal(t, 1, validatedResult)

	t.Run("Test to supply IBAN with spaces", func(t *testing.T) {
		validatedResult = validator.ValidateIban("GB82 WEST 1234569 8765432")
		assert.Equal(t, 1, validatedResult)

		validatedResult = validator.ValidateIban(" NI 92BA MC00 00000000 0000000312 3123 ")
		assert.Equal(t, 1, validatedResult)

		validatedResult = validator.ValidateIban(" NI 92BA MC00 00010000 0000000312 3123 ")
		assert.Equal(t, 0, validatedResult)
	})

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
