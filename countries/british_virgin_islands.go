package countries

import "github.com/rojack96/mcc/models"

var BritishVirginIslands = models.Mcc{
	Code:        348,
	Iso:         "VG",
	Country:     "British Virgin Islands",
	CountryCode: 284,
	Mnc: []models.Mnc{
		{Code: "170", Network: "LIME"},
		{Code: "570", Network: "Caribbean Cellular"},
		{Code: "770", Network: "Digicel"},
	},
}
