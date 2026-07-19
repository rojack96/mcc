package countries

import "github.com/rojack96/mcc/models"

var Croatia = models.Mcc{
	Code:        219,
	Iso:         "HR",
	Country:     "Croatia",
	CountryCode: 385,
	Mnc: []models.Mnc{
		{Code: "10", Network: "A1 / VIP"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "1", Network: "T-Mobile"},
		{Code: "12", Network: "TELE FOCUS"},
		{Code: "2", Network: "Telemach / Tele2"},
	},
}
