package countries

import "github.com/rojack96/mcc/models"

var Philippines = models.Mcc{
	Code:        515,
	Iso:         "PH",
	Country:     "Philippines",
	CountryCode: 63,
	Mnc: []models.Mnc{
		{Code: "02", Network: "Globe Telecom"},
		{Code: "03", Network: "Smart"},
		{Code: "05", Network: "SUN/Digitel"},
		{Code: "18", Network: "RED Mobile/Cure"},
	},
}
