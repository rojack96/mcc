package countries

import "github.com/rojack96/mcc/models"

var Mozambique = models.Mcc{
	Code:        643,
	Iso:         "MZ",
	Country:     "Mozambique",
	CountryCode: 258,
	Mnc: []models.Mnc{
		{Code: "01", Network: "mCel"},
		{Code: "03", Network: "Movitel"},
		{Code: "04", Network: "Vodacom Sarl"},
	},
}
