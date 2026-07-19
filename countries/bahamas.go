package countries

import "github.com/rojack96/mcc/models"

var Bahamas = models.Mcc{
	Code:        364,
	Iso:         "BS",
	Country:     "Bahamas",
	CountryCode: 1242,
	Mnc: []models.Mnc{
		{Code: "39", Network: "Bahamas Telco. Comp."},
		{Code: "390", Network: "Bahamas Telco. Comp."},
	},
}
