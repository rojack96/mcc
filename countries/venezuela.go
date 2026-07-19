package countries

import "github.com/rojack96/mcc/models"

var Venezuela = models.Mcc{
	Code:        734,
	Iso:         "VE",
	Country:     "Venezuela",
	CountryCode: 58,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Infonet"},
		{Code: "02", Network: "DigiTel C.A."},
		{Code: "03", Network: "Digicel"},
		{Code: "04", Network: "Movistar/TelCel"},
		{Code: "06", Network: "Movilnet C.A."},
	},
}
