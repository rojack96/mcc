package countries

import "github.com/rojack96/mcc/models"

var Andorra = models.Mcc{
	Code:        213,
	Iso:         "AD",
	Country:     "Andorra",
	CountryCode: 376,
	Mnc: []models.Mnc{
		{Code: "3", Network: "Andorra Telecom / Mobiland"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
	},
}
