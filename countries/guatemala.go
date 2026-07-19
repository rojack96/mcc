package countries

import "github.com/rojack96/mcc/models"

var Guatemala = models.Mcc{
	Code:        704,
	Iso:         "GT",
	Country:     "Guatemala",
	CountryCode: 502,
	Mnc: []models.Mnc{
		{Code: "01", Network: "SERCOM"},
		{Code: "02", Network: "TIGO/COMCEL"},
		{Code: "03", Network: "Telefonica"},
	},
}
