package countries

import "github.com/rojack96/mcc/models"

var CostaRica = models.Mcc{
	Code:        712,
	Iso:         "CR",
	Country:     "Costa Rica",
	CountryCode: 506,
	Mnc: []models.Mnc{
		{Code: "01", Network: "ICE"},
		{Code: "02", Network: "ICE"},
		{Code: "03", Network: "Claro"},
		{Code: "04", Network: "Movistar"},
		{Code: "999", Network: "Fix Line Costa Rica"},
	},
}
