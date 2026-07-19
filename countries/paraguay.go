package countries

import "github.com/rojack96/mcc/models"

var Paraguay = models.Mcc{
	Code:        744,
	Iso:         "PY",
	Country:     "Paraguay",
	CountryCode: 595,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Hola/VOX"},
		{Code: "02", Network: "Claro/Hutchison"},
		{Code: "04", Network: "Tigo/Telecel"},
		{Code: "05", Network: "Personal"},
	},
}
