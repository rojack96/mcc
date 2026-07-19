package countries

import "github.com/rojack96/mcc/models"

var Egypt = models.Mcc{
	Code:        602,
	Iso:         "EG",
	Country:     "Egypt",
	CountryCode: 20,
	Mnc: []models.Mnc{
		{Code: "01", Network: "EMS - Mobinil"},
		{Code: "02", Network: "Vodafone (Misrfone Telecom)"},
		{Code: "03", Network: "ETISALAT"},
		{Code: "04", Network: "WE/Telecom"},
	},
}
