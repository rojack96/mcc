package countries

import "github.com/rojack96/mcc/models"

var Suriname = models.Mcc{
	Code:        746,
	Iso:         "SR",
	Country:     "Suriname",
	CountryCode: 597,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Telesur"},
		{Code: "02", Network: "TELESUR 2"},
		{Code: "03", Network: "Digicel"},
		{Code: "04", Network: "UNIQA"},
	},
}
