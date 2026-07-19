package countries

import "github.com/rojack96/mcc/models"

var TimorLeste = models.Mcc{
	Code:        514,
	Iso:         "TL",
	Country:     "Timor-Leste",
	CountryCode: 670,
	Mnc: []models.Mnc{
		{Code: "02", Network: "Timor Telecom"},
		{Code: "999", Network: "Fix Line"},
	},
}
