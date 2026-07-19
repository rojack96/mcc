package countries

import "github.com/rojack96/mcc/models"

var Kuwait = models.Mcc{
	Code:        419,
	Iso:         "KW",
	Country:     "Kuwait",
	CountryCode: 965,
	Mnc: []models.Mnc{
		{Code: "02", Network: "Zain"},
		{Code: "03", Network: "STC"},
		{Code: "04", Network: "Viva"},
		{Code: "09", Network: "Virgin Mobile"},
	},
}
