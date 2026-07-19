package countries

import "github.com/rojack96/mcc/models"

var Seychelles = models.Mcc{
	Code:        633,
	Iso:         "SC",
	Country:     "Seychelles",
	CountryCode: 248,
	Mnc: []models.Mnc{
		{Code: "01", Network: "C&W"},
		{Code: "02", Network: "Smartcom"},
		{Code: "10", Network: "Airtel"},
	},
}
