package countries

import "github.com/rojack96/mcc/models"

var Morocco = models.Mcc{
	Code:        604,
	Iso:         "MA",
	Country:     "Morocco",
	CountryCode: 212,
	Mnc: []models.Mnc{
		{Code: "00", Network: "Orange/Medi Telecom"},
		{Code: "01", Network: "IAM/Itissallat"},
		{Code: "02", Network: "INWI/WANA"},
		{Code: "04", Network: "Al Houria Telecom"},
		{Code: "05", Network: "INWI/WANA"},
		{Code: "06", Network: "IAM/Itissallat"},
		{Code: "99", Network: "Al Houria Telecom"},
	},
}
