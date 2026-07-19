package countries

import "github.com/rojack96/mcc/models"

var FrenchPolynesia = models.Mcc{
	Code:        547,
	Iso:         "PF",
	Country:     "French Polynesia",
	CountryCode: 689,
	Mnc: []models.Mnc{
		{Code: "15", Network: "Pacific Mobile Telecom (PMT)"},
		{Code: "20", Network: "Tikiphone"},
	},
}
