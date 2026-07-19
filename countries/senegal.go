package countries

import "github.com/rojack96/mcc/models"

var Senegal = models.Mcc{
	Code:        608,
	Iso:         "SN",
	Country:     "Senegal",
	CountryCode: 221,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Orange/Sonatel"},
		{Code: "02", Network: "Sentel GSM"},
		{Code: "03", Network: "Expresso/Sudatel"},
	},
}
