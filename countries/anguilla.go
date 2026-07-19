package countries

import "github.com/rojack96/mcc/models"

var Anguilla = models.Mcc{
	Code:        365,
	Iso:         "AI",
	Country:     "Anguilla",
	CountryCode: 1264,
	Mnc: []models.Mnc{
		{Code: "010", Network: "Wireless Ventures Ltd"},
		{Code: "840", Network: "Cable and Wireless"},
	},
}
