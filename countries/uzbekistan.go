package countries

import "github.com/rojack96/mcc/models"

var Uzbekistan = models.Mcc{
	Code:        434,
	Iso:         "UZ",
	Country:     "Uzbekistan",
	CountryCode: 998,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Buztel"},
		{Code: "04", Network: "Bee Line/Unitel"},
		{Code: "05", Network: "Ucell/Coscom"},
		{Code: "07", Network: "MTS/Uzdunrobita"},
	},
}
