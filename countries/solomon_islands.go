package countries

import "github.com/rojack96/mcc/models"

var SolomonIslands = models.Mcc{
	Code:        540,
	Iso:         "SB",
	Country:     "Solomon Islands",
	CountryCode: 677,
	Mnc: []models.Mnc{
		{Code: "01", Network: "BREEZE"},
		{Code: "02", Network: "bemobile"},
	},
}
