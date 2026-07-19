package countries

import "github.com/rojack96/mcc/models"

var Djibouti = models.Mcc{
	Code:        638,
	Iso:         "DJ",
	Country:     "Djibouti",
	CountryCode: 253,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Djibouti Telecom SA (Evatis)"},
	},
}
