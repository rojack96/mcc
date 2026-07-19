package countries

import "github.com/rojack96/mcc/models"

var Aruba = models.Mcc{
	Code:        363,
	Iso:         "AW",
	Country:     "Aruba",
	CountryCode: 297,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Setar GSM"},
		{Code: "02", Network: "Digicel"},
		{Code: "20", Network: "Digicel"},
	},
}
