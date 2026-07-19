package countries

import "github.com/rojack96/mcc/models"

var Cuba = models.Mcc{
	Code:        368,
	Iso:         "CU",
	Country:     "Cuba",
	CountryCode: 53,
	Mnc: []models.Mnc{
		{Code: "01", Network: "C-COM"},
		{Code: "999", Network: "Fix Line Cuba"},
	},
}
