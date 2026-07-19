package countries

import "github.com/rojack96/mcc/models"

var Bahrain = models.Mcc{
	Code:        426,
	Iso:         "BH",
	Country:     "Bahrain",
	CountryCode: 973,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Batelco"},
		{Code: "02", Network: "MTC Vodafone"},
		{Code: "04", Network: "VIVA"},
	},
}
