package countries

import "github.com/rojack96/mcc/models"

var NewCaledonia = models.Mcc{
	Code:        546,
	Iso:         "NC",
	Country:     "New Caledonia",
	CountryCode: 687,
	Mnc: []models.Mnc{
		{Code: "01", Network: "OPT Mobilis"},
	},
}
