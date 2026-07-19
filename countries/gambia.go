package countries

import "github.com/rojack96/mcc/models"

var Gambia = models.Mcc{
	Code:        607,
	Iso:         "GM",
	Country:     "Gambia",
	CountryCode: 220,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Gamcel"},
		{Code: "02", Network: "Africel"},
		{Code: "03", Network: "Comium"},
		{Code: "04", Network: "Q-Cell"},
	},
}
