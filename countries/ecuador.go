package countries

import "github.com/rojack96/mcc/models"

var Ecuador = models.Mcc{
	Code:        740,
	Iso:         "EC",
	Country:     "Ecuador",
	CountryCode: 593,
	Mnc: []models.Mnc{
		{Code: "00", Network: "MOVISTAR/OteCel"},
		{Code: "01", Network: "Porta/Conecel"},
		{Code: "02", Network: "CNT Mobile"},
		{Code: "03", Network: "Tuenti"},
	},
}
