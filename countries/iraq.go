package countries

import "github.com/rojack96/mcc/models"

var Iraq = models.Mcc{
	Code:        418,
	Iso:         "IQ",
	Country:     "Iraq",
	CountryCode: 964,
	Mnc: []models.Mnc{
		{Code: "05", Network: "Asia Cell"},
		{Code: "20", Network: "Orascom Telecom"},
		{Code: "40", Network: "Korek"},
		{Code: "45", Network: "Mobitel (Iraq-Kurdistan) and Moutiny"},
		{Code: "66", Network: "Fastlink"},
		{Code: "82", Network: "Korek"},
	},
}
