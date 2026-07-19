package countries

import "github.com/rojack96/mcc/models"

var Cyprus = models.Mcc{
	Code:        280,
	Iso:         "CY",
	Country:     "Cyprus",
	CountryCode: 357,
	Mnc: []models.Mnc{
		{Code: "22", Network: "Cablenet / Lemontel"},
		{Code: "02", Network: "Cytamobile-Vodafone"},
		{Code: "01", Network: "Cytamobile-Vodafone"},
		{Code: "10", Network: "Epic / MTN"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "20", Network: "PrimeTel"},
	},
}
