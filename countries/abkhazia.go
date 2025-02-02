package countries

import "github.com/rojack96/mcc/models"

var Abkhazia = models.Mcc{
	Code:        289,
	Iso:         "GE",
	Country:     "Abkhazia",
	CountryCode: 7,
	Mnc: []models.Mnc{
		{Code: "88", Network: "A-Mobile"},
		{Code: "68", Network: "A-Mobile"},
		{Code: "67", Network: "Aquafon"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
	},
}
