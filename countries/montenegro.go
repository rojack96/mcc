package countries

import "github.com/rojack96/mcc/models"

var Montenegro = models.Mcc{
	Code:        297,
	Iso:         "ME",
	Country:     "Montenegro",
	CountryCode: 382,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "03", Network: "Mtel"},
		{Code: "02", Network: "Telekom / T-mobile"},
		{Code: "01", Network: "Telenor"},
	},
}
