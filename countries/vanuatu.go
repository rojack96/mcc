package countries

import "github.com/rojack96/mcc/models"

var Vanuatu = models.Mcc{
	Code:        541,
	Iso:         "VU",
	Country:     "Vanuatu",
	CountryCode: 678,
	Mnc: []models.Mnc{
		{Code: "01", Network: "SMILE"},
	},
}
