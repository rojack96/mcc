package countries

import "github.com/rojack96/mcc/models"

var Niue = models.Mcc{
	Code:        555,
	Iso:         "NU",
	Country:     "Niue",
	CountryCode: 683,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Niue Telecom"},
	},
}
