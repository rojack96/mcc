package countries

import "github.com/rojack96/mcc/models"

var Togo = models.Mcc{
	Code:        615,
	Iso:         "TG",
	Country:     "Togo",
	CountryCode: 228,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Togocom Yas"},
		{Code: "03", Network: "Telecel/MOOV"},
	},
}
