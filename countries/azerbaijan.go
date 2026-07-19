package countries

import "github.com/rojack96/mcc/models"

var Azerbaijan = models.Mcc{
	Code:        400,
	Iso:         "AZ",
	Country:     "Azerbaijan",
	CountryCode: 994,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Azercell Telekom B.M."},
		{Code: "02", Network: "J.V. Bakcell GSM 2000"},
		{Code: "03", Network: "CATEL"},
		{Code: "04", Network: "Azerfon."},
		{Code: "06", Network: "Naxtel"},
	},
}
