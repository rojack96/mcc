package countries

import "github.com/rojack96/mcc/models"

var Kyrgyzstan = models.Mcc{
	Code:        437,
	Iso:         "KG",
	Country:     "Kyrgyzstan",
	CountryCode: 996,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Beeline/Bitel"},
		{Code: "03", Network: "AkTel/Fonex LLC"},
		{Code: "05", Network: "MEGACOM"},
		{Code: "09", Network: "O!/NUR Telecom"},
	},
}
