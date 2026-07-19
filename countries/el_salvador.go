package countries

import "github.com/rojack96/mcc/models"

var ElSalvador = models.Mcc{
	Code:        706,
	Iso:         "SV",
	Country:     "El Salvador",
	CountryCode: 503,
	Mnc: []models.Mnc{
		{Code: "01", Network: "CLARO/CTE"},
		{Code: "02", Network: "Digicel"},
		{Code: "03", Network: "Telemovil"},
		{Code: "04", Network: "Telefonica"},
		{Code: "05", Network: "INTELFON SA de CV"},
	},
}
