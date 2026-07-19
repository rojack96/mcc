package countries

import "github.com/rojack96/mcc/models"

var PapuaNewGuinea = models.Mcc{
	Code:        537,
	Iso:         "PG",
	Country:     "Papua New Guinea",
	CountryCode: 675,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Telikom"},
		{Code: "02", Network: "Telikom"},
		{Code: "03", Network: "Digicel"},
	},
}
