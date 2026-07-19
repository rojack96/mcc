package countries

import "github.com/rojack96/mcc/models"

var Libya = models.Mcc{
	Code:        606,
	Iso:         "LY",
	Country:     "Libya",
	CountryCode: 218,
	Mnc: []models.Mnc{
		{Code: "00", Network: "Libyana Mobile Phone"},
		{Code: "01", Network: "Al-Madar"},
		{Code: "02", Network: "Al-Jeel"},
		{Code: "03", Network: "Libyana Mobile Phone"},
	},
}
