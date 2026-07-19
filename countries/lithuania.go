package countries

import "github.com/rojack96/mcc/models"

var Lithuania = models.Mcc{
	Code:        246,
	Iso:         "LT",
	Country:     "Lithuania",
	CountryCode: 370,
	Mnc: []models.Mnc{
		{Code: "02", Network: "Bite"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "05", Network: "LTG"},
		{Code: "06", Network: "Mediafon"},
		{Code: "299", Network: "SkyCall"},
		{Code: "03", Network: "Tele2"},
		{Code: "299", Network: "Teletel"},
		{Code: "01", Network: "Telia"},
	},
}
