package countries

import "github.com/rojack96/mcc/models"

var Bulgaria = models.Mcc{
	Code:        284,
	Iso:         "BG",
	Country:     "Bulgaria",
	CountryCode: 359,
	Mnc: []models.Mnc{
		{Code: "01", Network: "A1"},
		{Code: "11", Network: "Bulsatcom"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "13", Network: "T.com"},
		{Code: "05", Network: "Telenor"},
		{Code: "03", Network: "Vivacom"},
	},
}
