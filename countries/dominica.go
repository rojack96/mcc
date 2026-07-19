package countries

import "github.com/rojack96/mcc/models"

var Dominica = models.Mcc{
	Code:        366,
	Iso:         "DM",
	Country:     "Dominica",
	CountryCode: 1767,
	Mnc: []models.Mnc{
		{Code: "020", Network: "DigiCel / Cingular Wireless"},
		{Code: "110", Network: "C & W"},
	},
}
