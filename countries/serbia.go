package countries

import "github.com/rojack96/mcc/models"

var Serbia = models.Mcc{
	Code:        220,
	Iso:         "RS",
	Country:     "Serbia",
	CountryCode: 381,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "11", Network: "Globaltel"},
		{Code: "3", Network: "MTS"},
		{Code: "1", Network: "Telenor"},
		{Code: "5", Network: "VIP"},
		{Code: "20", Network: "VIP"},
	},
}
