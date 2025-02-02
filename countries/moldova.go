package countries

import "github.com/rojack96/mcc/models"

var Moldova = models.Mcc{
	Code:        259,
	Iso:         "MD",
	Country:     "Moldova",
	CountryCode: 373,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "02", Network: "Moldcell"},
		{Code: "01", Network: "Orange"},
		{Code: "03", Network: "Unite"},
		{Code: "99", Network: "Unite"},
		{Code: "05", Network: "Unite"},
	},
}
