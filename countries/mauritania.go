package countries

import "github.com/rojack96/mcc/models"

var Mauritania = models.Mcc{
	Code:        609,
	Iso:         "MR",
	Country:     "Mauritania",
	CountryCode: 222,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Mattel"},
		{Code: "02", Network: "Chinguitel SA"},
		{Code: "10", Network: "Mauritel"},
	},
}
