package countries

import "github.com/rojack96/mcc/models"

var Kazakhstan = models.Mcc{
	Code:        401,
	Iso:         "KZ",
	Country:     "Kazakhstan",
	CountryCode: 7,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Beeline/KaR-Tel LLP"},
		{Code: "02", Network: "K-Cell"},
		{Code: "07", Network: "Dalacom/Altel"},
		{Code: "77", Network: "MTS"},
	},
}
