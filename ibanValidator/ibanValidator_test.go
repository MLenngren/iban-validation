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

	// NI is taken from https://www.iban.com/structure experimental section.
	// It fails the length test since the register "https://www.swift.com/swift-resource/9606/download" do not have that.
	validatedResult = validator.ValidateIban("NI92BAMC000000000000000003123123")
	assert.Equal(t, 0, validatedResult)

	// I noticed lack of removal of spaces when pasting an IBAN from another site, so decided to add that
	t.Run("Test to supply IBAN with spaces", func(t *testing.T) {
		validatedResult = validator.ValidateIban("GB82 WEST 1234569 8765432")
		assert.Equal(t, 1, validatedResult)

		validatedResult = validator.ValidateIban(" MU43BOM M010112345 678910 1 000MUR ")
		assert.Equal(t, 1, validatedResult)

		validatedResult = validator.ValidateIban(" MU4 3BOMM0 10112345 6789101 000 MOR ")
		assert.Equal(t, 0, validatedResult)
	})

}

// According to some of the pages I read the first thing you should do to validate
// is to make sure the length of the iban is correct for the specific country
func TestGetIbanCountryLength(t *testing.T) {
	ibanLength := validator.GetIbanCountryLength("MU43BOMM0101123456789101000MUR")
	assert.Equal(t, 30, ibanLength)
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
