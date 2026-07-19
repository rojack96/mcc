package countries

import "github.com/rojack96/mcc/models"

var SyrianArabRepublic = models.Mcc{
	Code:        417,
	Iso:         "SY",
	Country:     "Syrian Arab Republic",
	CountryCode: 963,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Syriatel Mobile Telecom SA"},
		{Code: "02", Network: "MTN/Spacetel"},
	},
}
