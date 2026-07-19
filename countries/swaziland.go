package countries

import "github.com/rojack96/mcc/models"

var Swaziland = models.Mcc{
	Code:        653,
	Iso:         "SZ",
	Country:     "Swaziland",
	CountryCode: 268,
	Mnc: []models.Mnc{
		{Code: "01", Network: "SwaziTelecom"},
		{Code: "02", Network: "Swazi Mobile"},
		{Code: "10", Network: "Swazi MTN"},
	},
}
