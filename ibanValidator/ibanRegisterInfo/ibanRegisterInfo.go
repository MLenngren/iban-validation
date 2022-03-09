package ibanRegisterInfo

import (
	s "strings"
)

type ibanRegisterInfo struct {
	CountryCode string
	CountryName string
	IbanLength  int
}

// TODO: Create some kind of error, returning 0 if no country found is not correct
func GetIbanLengthForCountryCode(countryCode string) int {
	for _, ibanInfo := range ListOfIbanInfo {

		if ibanInfo.CountryCode == s.ToUpper(countryCode) {
			return ibanInfo.IbanLength
		}
	}
	return 0
}

// Iban list - retrieved and fixed by using a macro
// TODO: automate updating of below list and to DB.
var ListOfIbanInfo = []ibanRegisterInfo{
	{
		CountryCode: "AD",
		CountryName: "Andorra",
		IbanLength:  24,
	},
	{
		CountryCode: "AE",
		CountryName: "United Arab Emirates (The)",
		IbanLength:  23,
	},
	{
		CountryCode: "AL",
		CountryName: "Albania",
		IbanLength:  28,
	},
	{
		CountryCode: "AT",
		CountryName: "Austria",
		IbanLength:  20,
	},
	{
		CountryCode: "AZ",
		CountryName: "Azerbaijan",
		IbanLength:  28,
	},
	{
		CountryCode: "BA",
		CountryName: "Bosnia and Herzegovina",
		IbanLength:  20,
	},
	{
		CountryCode: "BE",
		CountryName: "Belgium",
		IbanLength:  16,
	},
	{
		CountryCode: "BG",
		CountryName: "Bulgaria",
		IbanLength:  22,
	},
	{
		CountryCode: "BH",
		CountryName: "Bahrain",
		IbanLength:  22,
	},
	{
		CountryCode: "BI",
		CountryName: "Burundi",
		IbanLength:  27,
	},
	{
		CountryCode: "BR",
		CountryName: "Brazil",
		IbanLength:  29,
	},
	{
		CountryCode: "BY",
		CountryName: "Republic of Belarus",
		IbanLength:  28,
	},
	{
		CountryCode: "CH",
		CountryName: "Switzerland",
		IbanLength:  21,
	},
	{
		CountryCode: "CR",
		CountryName: "Costa Rica",
		IbanLength:  22,
	},
	{
		CountryCode: "CY",
		CountryName: "Cyprus",
		IbanLength:  28,
	},
	{
		CountryCode: "CZ",
		CountryName: "Czech Republic",
		IbanLength:  24,
	},
	{
		CountryCode: "DE",
		CountryName: "Germany",
		IbanLength:  22,
	},
	{
		CountryCode: "DK",
		CountryName: "Denmark",
		IbanLength:  18,
	},
	{
		CountryCode: "DO",
		CountryName: "Dominican Republic",
		IbanLength:  28,
	},
	{
		CountryCode: "EE",
		CountryName: "Estonia",
		IbanLength:  20,
	},
	{
		CountryCode: "EG",
		CountryName: "Egypt",
		IbanLength:  29,
	},
	{
		CountryCode: "ES",
		CountryName: "Spain",
		IbanLength:  24,
	},
	{
		CountryCode: "FI",
		CountryName: "Finland",
		IbanLength:  18,
	},
	{
		CountryCode: "FO",
		CountryName: "Faroe Islands",
		IbanLength:  18,
	},
	{
		CountryCode: "FR",
		CountryName: "France",
		IbanLength:  27,
	},
	{
		CountryCode: "GB",
		CountryName: "United Kingdom",
		IbanLength:  22,
	},
	{
		CountryCode: "GE",
		CountryName: "Georgia",
		IbanLength:  22,
	},
	{
		CountryCode: "GI",
		CountryName: "Gibraltar",
		IbanLength:  23,
	},
	{
		CountryCode: "GL",
		CountryName: "Greenland",
		IbanLength:  18,
	},
	{
		CountryCode: "GR",
		CountryName: "Greece",
		IbanLength:  27,
	},
	{
		CountryCode: "GT",
		CountryName: "Guatemala",
		IbanLength:  28,
	},
	{
		CountryCode: "HR",
		CountryName: "Croatia",
		IbanLength:  21,
	},
	{
		CountryCode: "HU",
		CountryName: "Hungary",
		IbanLength:  28,
	},
	{
		CountryCode: "IE",
		CountryName: "Ireland",
		IbanLength:  22,
	},
	{
		CountryCode: "IL",
		CountryName: "Israel",
		IbanLength:  23,
	},
	{
		CountryCode: "IQ",
		CountryName: "Iraq",
		IbanLength:  23,
	},
	{
		CountryCode: "IS",
		CountryName: "Iceland",
		IbanLength:  26,
	},
	{
		CountryCode: "IT",
		CountryName: "Italy",
		IbanLength:  27,
	},
	{
		CountryCode: "JO",
		CountryName: "Jordan",
		IbanLength:  30,
	},
	{
		CountryCode: "KW",
		CountryName: "Kuwait",
		IbanLength:  30,
	},
	{
		CountryCode: "KZ",
		CountryName: "Kazakhstan",
		IbanLength:  20,
	},
	{
		CountryCode: "LB",
		CountryName: "Lebanon",
		IbanLength:  28,
	},
	{
		CountryCode: "LC",
		CountryName: "Saint Lucia",
		IbanLength:  32,
	},
	{
		CountryCode: "LI",
		CountryName: "Liechtenstein",
		IbanLength:  21,
	},
	{
		CountryCode: "LT",
		CountryName: "Lithuania",
		IbanLength:  20,
	},
	{
		CountryCode: "LU",
		CountryName: "Luxembourg",
		IbanLength:  20,
	},
	{
		CountryCode: "LV",
		CountryName: "Latvia",
		IbanLength:  21,
	},
	{
		CountryCode: "LY",
		CountryName: "Libya",
		IbanLength:  25,
	},
	{
		CountryCode: "MC",
		CountryName: "Monaco",
		IbanLength:  27,
	},
	{
		CountryCode: "MD",
		CountryName: "Moldova",
		IbanLength:  24,
	},
	{
		CountryCode: "ME",
		CountryName: "Montenegro",
		IbanLength:  22,
	},
	{
		CountryCode: "MK",
		CountryName: "Macedonia",
		IbanLength:  19,
	},
	{
		CountryCode: "MR",
		CountryName: "Mauritania",
		IbanLength:  27,
	},
	{
		CountryCode: "MT",
		CountryName: "Malta",
		IbanLength:  31,
	},
	{
		CountryCode: "MU",
		CountryName: "Mauritius",
		IbanLength:  30,
	},
	{
		CountryCode: "NL",
		CountryName: "Netherlands (The)",
		IbanLength:  18,
	},
	{
		CountryCode: "NO",
		CountryName: "Norway",
		IbanLength:  15,
	},
	{
		CountryCode: "PK",
		CountryName: "Pakistan",
		IbanLength:  24,
	},
	{
		CountryCode: "PL",
		CountryName: "Poland",
		IbanLength:  28,
	},
	{
		CountryCode: "PS",
		CountryName: "Palestine, State of",
		IbanLength:  29,
	},
	{
		CountryCode: "PT",
		CountryName: "Portugal",
		IbanLength:  25,
	},
	{
		CountryCode: "QA",
		CountryName: "Qatar",
		IbanLength:  29,
	},
	{
		CountryCode: "RO",
		CountryName: "Romania",
		IbanLength:  24,
	},
	{
		CountryCode: "RS",
		CountryName: "Serbia",
		IbanLength:  22,
	},
	{
		CountryCode: "SA",
		CountryName: "Saudi Arabia",
		IbanLength:  24,
	},
	{
		CountryCode: "SC",
		CountryName: "Seychelles",
		IbanLength:  31,
	},
	{
		CountryCode: "SD",
		CountryName: "Sudan",
		IbanLength:  18,
	},
	{
		CountryCode: "SE",
		CountryName: "Sweden",
		IbanLength:  24,
	},
	{
		CountryCode: "SI",
		CountryName: "Slovenia",
		IbanLength:  19,
	},
	{
		CountryCode: "SK",
		CountryName: "Slovakia",
		IbanLength:  24,
	},
	{
		CountryCode: "SM",
		CountryName: "San Marino",
		IbanLength:  27,
	},
	{
		CountryCode: "ST",
		CountryName: "Sao Tome and Principe",
		IbanLength:  25,
	},
	{
		CountryCode: "SV",
		CountryName: "El Salvador",
		IbanLength:  28,
	},
	{
		CountryCode: "TL",
		CountryName: "Timor-Leste",
		IbanLength:  23,
	},
	{
		CountryCode: "TN",
		CountryName: "Tunisia",
		IbanLength:  24,
	},
	{
		CountryCode: "TR",
		CountryName: "Turkey",
		IbanLength:  26,
	},
	{
		CountryCode: "UA",
		CountryName: "Ukraine",
		IbanLength:  29,
	},
	{
		CountryCode: "VA",
		CountryName: "Vatican City State",
		IbanLength:  22,
	},
	{
		CountryCode: "VG",
		CountryName: "Virgin Islands",
		IbanLength:  24,
	},
	{
		CountryCode: "XK",
		CountryName: "Kosovo",
		IbanLength:  20,
	},
}
