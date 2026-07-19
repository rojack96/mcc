package countries

import "github.com/rojack96/mcc/models"

var Chad = models.Mcc{
	Code:        622,
	Iso:         "TD",
	Country:     "Chad",
	CountryCode: 235,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Zain/Airtel/Celtel"},
		{Code: "03", Network: "Tigo/Milicom"},
		{Code: "04", Network: "Salam/Sotel"},
	},
}
