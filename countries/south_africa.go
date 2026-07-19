package countries

import "github.com/rojack96/mcc/models"

var SouthAfrica = models.Mcc{
	Code:        655,
	Iso:         "ZA",
	Country:     "South Africa",
	CountryCode: 27,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Vodacom"},
		{Code: "02", Network: "8.ta"},
		{Code: "07", Network: "Cell C"},
		{Code: "10", Network: "MTN"},
		{Code: "19", Network: "Wireless Business Solutions (Pty) Ltd"},
	},
}
