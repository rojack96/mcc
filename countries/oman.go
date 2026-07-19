package countries

import "github.com/rojack96/mcc/models"

var Oman = models.Mcc{
	Code:        422,
	Iso:         "OM",
	Country:     "Oman",
	CountryCode: 968,
	Mnc: []models.Mnc{
		{Code: "02", Network: "Oman Mobile/GTO"},
		{Code: "03", Network: "Nawras"},
	},
}
