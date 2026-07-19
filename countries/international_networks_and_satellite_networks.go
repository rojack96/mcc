package countries

import "github.com/rojack96/mcc/models"

var InternationalNetworksAndSatelliteNetworks = models.Mcc{
	Code:        901,
	Iso:         "N/A",
	Country:     "International Networks and Satellite Networks",
	CountryCode: 0,
	Mnc: []models.Mnc{
		{Code: "05", Network: "Thuraya Satellite"},
		{Code: "11", Network: "InMarSAT"},
		{Code: "12", Network: "Maritime Communications Partner AS"},
		{Code: "13", Network: "Antarctica"},
		{Code: "14", Network: "AeroMobile"},
	},
}
