package countries

import "github.com/rojack96/mcc/models"

var Nepal = models.Mcc{
	Code:        429,
	Iso:         "NP",
	Country:     "Nepal",
	CountryCode: 977,
	Mnc: []models.Mnc{
		{Code: "01", Network: "NT Mobile / Namaste"},
		{Code: "02", Network: "Ncell"},
		{Code: "04", Network: "Smart Cell"},
		{Code: "999", Network: "Fix Line Nepal"},
	},
}
