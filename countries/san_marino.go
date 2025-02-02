package countries

import "github.com/rojack96/mcc/models"

var SanMarino = models.Mcc{
	Code:        292,
	Iso:         "SM",
	Country:     "San Marino",
	CountryCode: 378,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "01", Network: "Prima"},
		{Code: "299", Network: "TeleneT"},
	},
}
