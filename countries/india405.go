package countries

import "github.com/rojack96/mcc/models"

var India405 = models.Mcc{
	Code:        405,
	Iso:         "IN",
	Country:     "India",
	CountryCode: 91,
	Mnc: []models.Mnc{
		{Code: "034", Network: "TATA / Karnataka"},
		{Code: "05", Network: "Fascel Limited"},
		{Code: "87", Network: "Reliance Telecom Private Ltd."},
	},
}
