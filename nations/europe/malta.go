package europe

import "github.com/rojack96/mcc/models"

var Malta = models.Mcc{
	Code:        278,
	Iso:         "MT",
	Country:     "Malta",
	CountryCode: 356,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Epic"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "21", Network: "GO Mobile"},
		{Code: "30", Network: "GO Mobile"},
		{Code: "77", Network: "Melita"},
	},
}
