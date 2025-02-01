package europe

import "github.com/rojack96/mcc/models"

var Kosovo = models.Mcc{
	Code:        221,
	Iso:         "XK",
	Country:     "Kosovo",
	CountryCode: 383,
	Mnc: []models.Mnc{
		{Code: "7", Network: "D3 mobile"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "2", Network: "IPKO"},
		{Code: "299", Network: "MTS"},
		{Code: "1", Network: "Vala"},
	},
}
