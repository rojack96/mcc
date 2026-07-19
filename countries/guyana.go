package countries

import "github.com/rojack96/mcc/models"

var Guyana = models.Mcc{
	Code:        738,
	Iso:         "GY",
	Country:     "Guyana",
	CountryCode: 592,
	Mnc: []models.Mnc{
		{Code: "01", Network: "DigiCel"},
		{Code: "02", Network: "Cellink Plus"},
	},
}
