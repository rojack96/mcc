package countries

import "github.com/rojack96/mcc/models"

var Monaco = models.Mcc{
	Code:        212,
	Iso:         "MC",
	Country:     "Monaco",
	CountryCode: 377,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "1", Network: "Monaco Telecom"},
		{Code: "10", Network: "Monaco Telecom"},
	},
}
