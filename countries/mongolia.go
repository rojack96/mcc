package countries

import "github.com/rojack96/mcc/models"

var Mongolia = models.Mcc{
	Code:        428,
	Iso:         "MN",
	Country:     "Mongolia",
	CountryCode: 976,
	Mnc: []models.Mnc{
		{Code: "88", Network: "Unitel"},
		{Code: "91", Network: "Skytel Co. Ltd"},
		{Code: "98", Network: "G-Mobile Corporation Ltd"},
		{Code: "99", Network: "Mobicom"},
	},
}
