package countries

import "github.com/rojack96/mcc/models"

var BurkinaFaso = models.Mcc{
	Code:        613,
	Iso:         "BF",
	Country:     "Burkina Faso",
	CountryCode: 226,
	Mnc: []models.Mnc{
		{Code: "01", Network: "TeleMob-OnaTel"},
		{Code: "02", Network: "ZAIN - CelTel"},
		{Code: "03", Network: "TeleCel"},
	},
}
