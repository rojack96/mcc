package countries

import "github.com/rojack96/mcc/models"

var Grenada = models.Mcc{
	Code:        352,
	Iso:         "GD",
	Country:     "Grenada",
	CountryCode: 1473,
	Mnc: []models.Mnc{
		{Code: "030", Network: "Digicel"},
		{Code: "050", Network: "Digicel Grenada Ltd"},
		{Code: "110", Network: "Cable & Wireless"},
	},
}
