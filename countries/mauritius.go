package countries

import "github.com/rojack96/mcc/models"

var Mauritius = models.Mcc{
	Code:        617,
	Iso:         "MU",
	Country:     "Mauritius",
	CountryCode: 230,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Orange/Cellplus"},
		{Code: "02", Network: "Mahanagar Telephone"},
		{Code: "10", Network: "Emtel Ltd"},
	},
}
