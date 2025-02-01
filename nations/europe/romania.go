package europe

import "github.com/rojack96/mcc/models"

var Romania = models.Mcc{
	Code:        226,
	Iso:         "RO",
	Country:     "Romania",
	CountryCode: 40,
	Mnc: []models.Mnc{
		{Code: "5", Network: "Digi Mobil"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "299", Network: "Iristel"},
		{Code: "16", Network: "Lycamobile"},
		{Code: "10", Network: "Orange"},
		{Code: "2", Network: "Telekom"},
		{Code: "3", Network: "Telekom"},
		{Code: "1", Network: "Vodafone"},
	},
}
