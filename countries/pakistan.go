package countries

import "github.com/rojack96/mcc/models"

var Pakistan = models.Mcc{
	Code:        410,
	Iso:         "PK",
	Country:     "Pakistan",
	CountryCode: 92,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Mobilink"},
		{Code: "03", Network: "UFONE/PAKTel"},
		{Code: "04", Network: "ZONG/CMPak"},
		{Code: "06", Network: "Telenor"},
		{Code: "07", Network: "Warid Telecom"},
	},
}
