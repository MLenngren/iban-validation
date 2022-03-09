package ibanValidator

import (
	"fmt"
	s "strings"
)

// argument iban is required to be an ASCII string
func ValidateIban(iban string) int {

	modifiedIban := RotateFirst4Chars(iban)
	modifiedIban = CharacterToNumbers(modifiedIban)
	return -2
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
