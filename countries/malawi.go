package countries

import "github.com/rojack96/mcc/models"

var Malawi = models.Mcc{
	Code:        650,
	Iso:         "MW",
	Country:     "Malawi",
	CountryCode: 265,
	Mnc: []models.Mnc{
		{Code: "01", Network: "TNM/Telekom Network Ltd."},
		{Code: "10", Network: "Zain/Celtel ltd."},
	},
}
