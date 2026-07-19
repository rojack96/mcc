package countries

import "github.com/rojack96/mcc/models"

var CaymanIslands = models.Mcc{
	Code:        346,
	Iso:         "KY",
	Country:     "Cayman Islands",
	CountryCode: 1345,
	Mnc: []models.Mnc{
		{Code: "006", Network: "Digicel Cayman Ltd."},
		{Code: "050", Network: "Digicel Cayman Ltd"},
		{Code: "140", Network: "LIME / Cable & Wirel."},
	},
}
