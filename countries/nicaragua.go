package countries

import "github.com/rojack96/mcc/models"

var Nicaragua = models.Mcc{
	Code:        710,
	Iso:         "NI",
	Country:     "Nicaragua",
	CountryCode: 505,
	Mnc: []models.Mnc{
		{Code: "30", Network: "Movistar"},
		{Code: "73", Network: "SERCOM S.A."},
		{Code: "999", Network: "Fix Line"},
	},
}
