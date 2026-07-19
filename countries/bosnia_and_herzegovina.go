package countries

import "github.com/rojack96/mcc/models"

var BosniaAndHerzegovina = models.Mcc{
	Code:        218,
	Iso:         "BA",
	Country:     "Bosnia and Herzegovina",
	CountryCode: 387,
	Mnc: []models.Mnc{
		{Code: "90", Network: "BH Mobile"},
		{Code: "3", Network: "Eronet"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "5", Network: "m:tel"},
	},
}
