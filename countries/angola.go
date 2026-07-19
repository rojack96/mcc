package countries

import "github.com/rojack96/mcc/models"

var Angola = models.Mcc{
	Code:        631,
	Iso:         "AO",
	Country:     "Angola",
	CountryCode: 244,
	Mnc: []models.Mnc{
		{Code: "02", Network: "Unitel"},
		{Code: "04", Network: "MoviCel"},
	},
}
