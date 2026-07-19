package countries

import "github.com/rojack96/mcc/models"

var Turkmenistan = models.Mcc{
	Code:        438,
	Iso:         "TM",
	Country:     "Turkmenistan",
	CountryCode: 993,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Barash Communication"},
		{Code: "02", Network: "TM-Cell"},
	},
}
