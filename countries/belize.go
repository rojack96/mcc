package countries

import "github.com/rojack96/mcc/models"

var Belize = models.Mcc{
	Code:        702,
	Iso:         "BZ",
	Country:     "Belize",
	CountryCode: 501,
	Mnc: []models.Mnc{
		{Code: "67", Network: "DigiCell"},
	},
}
