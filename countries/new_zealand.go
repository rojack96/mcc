package countries

import "github.com/rojack96/mcc/models"

var NewZealand = models.Mcc{
	Code:        530,
	Iso:         "NZ",
	Country:     "New Zealand",
	CountryCode: 64,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Vodafone"},
		{Code: "02", Network: "Telecom New Zealand CDMA Mobile Network"},
		{Code: "05", Network: "Telecom Mobile Ltd"},
		{Code: "24", Network: "Two Degrees Mobile Ltd"},
		{Code: "999", Network: "Fix Line"},
	},
}
