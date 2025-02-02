package europe

import "github.com/rojack96/mcc/models"

var Ukraine = models.Mcc{
	Code:        255,
	Iso:         "UA",
	Country:     "Ukraine",
	CountryCode: 380,
	Mnc: []models.Mnc{
		{Code: "07", Network: "3Mob"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "04", Network: "IT"},
		{Code: "02", Network: "Kyivstar"},
		{Code: "03", Network: "Kyivstar"},
		{Code: "06", Network: "lifecell"},
		{Code: "21", Network: "PEOPLEnet"},
		{Code: "99", Network: "Phoenix"},
		{Code: "01", Network: "Vodafone"},
	},
}
