package countries

import "github.com/rojack96/mcc/models"

var Maldives = models.Mcc{
	Code:        472,
	Iso:         "MV",
	Country:     "Maldives",
	CountryCode: 960,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Dhiraagu/C&W"},
		{Code: "02", Network: "Wataniya/WMOBILE"},
	},
}
