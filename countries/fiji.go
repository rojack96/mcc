package countries

import "github.com/rojack96/mcc/models"

var Fiji = models.Mcc{
	Code:        542,
	Iso:         "FJ",
	Country:     "Fiji",
	CountryCode: 679,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Vodafone"},
		{Code: "02", Network: "DigiCell"},
	},
}
