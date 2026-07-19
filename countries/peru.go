package countries

import "github.com/rojack96/mcc/models"

var Peru = models.Mcc{
	Code:        716,
	Iso:         "PE",
	Country:     "Peru",
	CountryCode: 51,
	Mnc: []models.Mnc{
		{Code: "06", Network: "Movistar"},
		{Code: "07", Network: "Nextel"},
		{Code: "10", Network: "Claro /Amer.Mov./TIM"},
		{Code: "15", Network: "Viettel Mobile"},
		{Code: "17", Network: "Nextel"},
	},
}
