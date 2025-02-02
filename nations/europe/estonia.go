package europe

import "github.com/rojack96/mcc/models"

var Estonia = models.Mcc{
	Code:        248,
	Iso:         "EE",
	Country:     "Estonia",
	CountryCode: 372,
	Mnc: []models.Mnc{
		{Code: "02", Network: "Elisa"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "03", Network: "Tele2"},
		{Code: "13", Network: "Telia"},
		{Code: "01", Network: "Telia"},
		{Code: "04", Network: "TravelSim"},
	},
}
