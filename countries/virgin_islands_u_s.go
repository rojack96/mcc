package countries

import "github.com/rojack96/mcc/models"

var VirginIslandsUS = models.Mcc{
	Code:        376,
	Iso:         "VI",
	Country:     "Virgin Islands, U.S.",
	CountryCode: 1340,
	Mnc: []models.Mnc{
		{Code: "50", Network: "Digicel"},
	},
}
