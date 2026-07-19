package countries

import "github.com/rojack96/mcc/models"

var Yemen = models.Mcc{
	Code:        421,
	Iso:         "YE",
	Country:     "Yemen",
	CountryCode: 967,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Sabaphone"},
		{Code: "02", Network: "MTN/Spacetel"},
		{Code: "03", Network: "Yemen Mob. CDMA"},
		{Code: "04", Network: "HITS/Y Unitel"},
	},
}
