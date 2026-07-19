package countries

import "github.com/rojack96/mcc/models"

var Haiti = models.Mcc{
	Code:        372,
	Iso:         "HT",
	Country:     "Haiti",
	CountryCode: 509,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Comcel"},
		{Code: "02", Network: "Digicel"},
		{Code: "03", Network: "National Telecom SA (NatCom)"},
	},
}
