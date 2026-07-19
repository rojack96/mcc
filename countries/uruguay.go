package countries

import "github.com/rojack96/mcc/models"

var Uruguay = models.Mcc{
	Code:        748,
	Iso:         "UY",
	Country:     "Uruguay",
	CountryCode: 598,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Ancel/Antel"},
		{Code: "07", Network: "MOVISTAR"},
		{Code: "10", Network: "Claro/AM Wireless"},
	},
}
