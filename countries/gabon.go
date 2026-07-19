package countries

import "github.com/rojack96/mcc/models"

var Gabon = models.Mcc{
	Code:        628,
	Iso:         "GA",
	Country:     "Gabon",
	CountryCode: 241,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Libertis S.A."},
		{Code: "02", Network: "MOOV/Telecel"},
		{Code: "03", Network: "ZAIN/Celtel Gabon S.A."},
		{Code: "04", Network: "Azur/Usan S.A."},
	},
}
