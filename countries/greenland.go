package countries

import "github.com/rojack96/mcc/models"

var Greenland = models.Mcc{
	Code:        290,
	Iso:         "GL",
	Country:     "Greenland",
	CountryCode: 299,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "01", Network: "Tele Greenland"},
	},
}
