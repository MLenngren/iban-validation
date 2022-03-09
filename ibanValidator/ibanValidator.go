package ibanValidator

import (
	"fmt"
	"log"
	"strconv"
	s "strings"
	"unicode"

	"github.com/mlenngren/iban-validator/ibanValidator/ibanRegisterInfo"
)

// argument iban is required to be an ASCII string
func ValidateIban(iban string) int {

	iban = SpaceStringsBuilder(iban) // Remove spaces

	// Check length against country specified lengths
	if GetIbanCountryLength(iban) != len(iban) {
		log.Printf("IBAN not corrent length")
		return 0
	}

	modifiedIban := RotateFirst4Chars(iban)
	modifiedIban = CharacterToNumbers(modifiedIban)
	return LargeMod97Calc(modifiedIban)

}

// Remove all the spaces in the string
func SpaceStringsBuilder(str string) string {
	var b s.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

// Get the correct iban length of the country this iban belongs to
func GetIbanCountryLength(iban string) int {
	return ibanRegisterInfo.GetIbanLengthForCountryCode(GetCountryCode(iban))
}

// first 2 characters of IBAN is the country code
func GetCountryCode(iban string) string {
	ibanCountry := substr(iban, 0, 2)
	return ibanCountry
}

// Calculate mod97 from a long number in string representation
/*
Found on Math.Stack
A nine digit number is formed by taking the leftmost 9 digits of x.
The mod of this number with respect to 97, r is obtained.
Then another nine digit number, q is formed by concatenating r and the next 7 digits of the number.
This process is continued till the last value of qmod97 is obtained.
If it is 1 then that validates the number .
*/
func LargeMod97Calc(ibanNumbersInString string) int {

	// Some IBANS are really large, we can't handle it so we have to do it in chunks
	tempNumberString := ""
	step := 0
	lastMod97 := 0
	for _, ch := range ibanNumbersInString {

		// first step is using the first 9 numbers
		if step == 0 {
			tempNumberString += string(ch)
			if len(tempNumberString) == 9 {
				intNumber, err := strconv.Atoi(tempNumberString) // convert the string to an integer so we can do the mod97
				if err != nil {
					log.Printf("error: %v", err.Error())
					return 0
				}

				lastMod97 = intNumber % 97

				tempNumberString = "" // Clear the number string, we have used it
				step = 1              // finished with first 9

			}
		} else {
			// subsequent steps uses the next 7 numbers and prefixes that with the mod97 number
			tempNumberString += string(ch)
			if len(tempNumberString) == 7 {
				mod97String := fmt.Sprint(lastMod97)              // Need the number as a string
				tempNumberString = mod97String + tempNumberString // prefix with mod97
				intNumber, err := strconv.Atoi(tempNumberString)  // convert the string to an integer so we can do the mod97
				if err != nil {
					log.Printf("error: %v", err.Error())
					return 0
				}

				lastMod97 = intNumber % 97
				tempNumberString = "" // Clear the number string, we have used it

			}
		}
	}

	// When we dont have more characters to process, we still need to process the ones that are in the temp number string
	if len(tempNumberString) > 0 {
		mod97String := fmt.Sprint(lastMod97)
		tempNumberString = mod97String + tempNumberString
		intNumber, err := strconv.Atoi(tempNumberString)
		if err != nil {
			log.Printf("error: %v", err.Error())
			return 0
		}
		lastMod97 = intNumber % 97
	}

	// if the last mod97 we made was 1, the IBAN is considered to be validated
	if lastMod97 == 1 {
		return 1
	}

	return 0
}

// Fix Iban characters, we only want numbers in our string
func CharacterToNumbers(ibanFirst4ToEnd string) string {
	modifiedIban := ""

	for _, ibanChar := range s.ToUpper(ibanFirst4ToEnd) {
		ibanCharInt := GetIbanCharInt(ibanChar)
		if ibanCharInt >= 48 && ibanCharInt <= 57 {
			modifiedIban += string(ibanChar)
		} else {
			modifiedIban += fmt.Sprint((ibanCharInt))
		}
	}
	return modifiedIban
}

// Get iban char int, uppercase is decreased by 55
func GetIbanCharInt(ibanChar rune) int {
	ibanCharInt := int(ibanChar)

	if ibanCharInt > 64 {
		ibanCharInt -= 55
	}

	return ibanCharInt
}

// Move the first 4 characters and put them in the same order but last
func RotateFirst4Chars(iban string) string {
	first4 := substr(iban, 0, 4)
	last := substr(iban, 4, len(iban)-4)
	ibanFirst4ToEnd := last + first4
	return ibanFirst4ToEnd
}

// Get a substring
// TODO: Validate if this added complexity helps with UTF8
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
