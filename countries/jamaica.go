package countries

import "github.com/rojack96/mcc/models"

var Jamaica = models.Mcc{
	Code:        338,
	Iso:         "JM",
	Country:     "Jamaica",
	CountryCode: 1876,
	Mnc: []models.Mnc{
		{Code: "020", Network: "Cable & Wireless"},
		{Code: "050", Network: "JM DIGICEL"},
		{Code: "180", Network: "Cable & Wireless Jamaica Ltd."},
	},
}
