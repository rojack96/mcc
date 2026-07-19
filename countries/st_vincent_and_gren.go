package countries

import "github.com/rojack96/mcc/models"

var StVincentAndGren = models.Mcc{
	Code:        360,
	Iso:         "VC",
	Country:     "St. Vincent & Gren.",
	CountryCode: 1784,
	Mnc: []models.Mnc{
		{Code: "100", Network: "Cingular Wireless"},
		{Code: "110", Network: "C & W"},
		{Code: "70", Network: "Digicel (St. Vincent and Grenadines) Limited"},
	},
}
