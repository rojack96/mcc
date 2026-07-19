package countries

import "github.com/rojack96/mcc/models"

var Rwanda = models.Mcc{
	Code:        635,
	Iso:         "RW",
	Country:     "Rwanda",
	CountryCode: 250,
	Mnc: []models.Mnc{
		{Code: "10", Network: "MTN/Rwandacell"},
		{Code: "13", Network: "TIGO"},
		{Code: "14", Network: "Airtel Rwanda Ltd"},
	},
}
