package countries

import "github.com/rojack96/mcc/models"

var Myanmar = models.Mcc{
	Code:        414,
	Iso:         "MM",
	Country:     "Myanmar",
	CountryCode: 95,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Myanmar Post & Teleco."},
		{Code: "05", Network: "Oreedoo"},
		{Code: "06", Network: "Telenor"},
		{Code: "09", Network: "Mytel"},
		{Code: "999", Network: "Fix Line"},
	},
}
