package countries

import "github.com/rojack96/mcc/models"

var NorthMacedonia = models.Mcc{
	Code:        294,
	Iso:         "MK",
	Country:     "North Macedonia",
	CountryCode: 389,
	Mnc: []models.Mnc{
		{Code: "03", Network: "A1"},
		{Code: "02", Network: "A1"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "299", Network: "LATRON"},
		{Code: "04", Network: "Lycamobile"},
		{Code: "11", Network: "Mobik"},
		{Code: "299", Network: "Telekabel"},
		{Code: "01", Network: "Telekom"},
	},
}
