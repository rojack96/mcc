package countries

import "github.com/rojack96/mcc/models"

var Madagascar = models.Mcc{
	Code:        646,
	Iso:         "MG",
	Country:     "Madagascar",
	CountryCode: 261,
	Mnc: []models.Mnc{
		{Code: "01", Network: "MADACOM"},
		{Code: "02", Network: "Orange/Soci"},
		{Code: "03", Network: "Sacel"},
		{Code: "04", Network: "Telma"},
	},
}
