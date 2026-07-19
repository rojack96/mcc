package countries

import "github.com/rojack96/mcc/models"

var SaintKittsAndNevis = models.Mcc{
	Code:        356,
	Iso:         "KN",
	Country:     "Saint Kitts and Nevis",
	CountryCode: 1869,
	Mnc: []models.Mnc{
		{Code: "110", Network: "Cable & Wireless"},
		{Code: "50", Network: "Digicel"},
		{Code: "70", Network: "UTS Cariglobe"},
	},
}
