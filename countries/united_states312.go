package countries

import "github.com/rojack96/mcc/models"

var UnitedStates312 = models.Mcc{
	Code:        312,
	Iso:         "US",
	Country:     "United States",
	CountryCode: 1,
	Mnc: []models.Mnc{
		{Code: "030", Network: "Cross Wireless Telephone Co."},
		{Code: "040", Network: "Custer Telephone Cooperative Inc."},
		{Code: "090", Network: "Allied Wireless Communications Corporation"},
	},
}
