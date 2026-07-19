package countries

import "github.com/rojack96/mcc/models"

var Tonga = models.Mcc{
	Code:        539,
	Iso:         "TO",
	Country:     "Tonga",
	CountryCode: 676,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Tonga Communications Corporation"},
		{Code: "43", Network: "Shoreline Communication"},
		{Code: "999", Network: "Fix Line"},
	},
}
