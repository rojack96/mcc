package countries

import "github.com/rojack96/mcc/models"

var Tajikistan = models.Mcc{
	Code:        436,
	Iso:         "TJ",
	Country:     "Tajikistan",
	CountryCode: 992,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Tcell/JC Somoncom"},
		{Code: "02", Network: "CJSC Indigo Tajikistan"},
		{Code: "03", Network: "Megafon"},
		{Code: "04", Network: "Babilon-M"},
		{Code: "05", Network: "Bee Line"},
		{Code: "12", Network: "Indigo Tajikistan"},
	},
}
