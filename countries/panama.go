package countries

import "github.com/rojack96/mcc/models"

var Panama = models.Mcc{
	Code:        714,
	Iso:         "PA",
	Country:     "Panama",
	CountryCode: 507,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Cable & Wireless S.A."},
		{Code: "02", Network: "Movistar"},
		{Code: "020", Network: "Movistar"},
		{Code: "03", Network: "Claro"},
		{Code: "04", Network: "Digicel"},
		{Code: "999", Network: "Fix Line"},
	},
}
