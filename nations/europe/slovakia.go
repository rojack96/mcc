package europe

import "github.com/rojack96/mcc/models"

var Slovakia = models.Mcc{
	Code:        231,
	Iso:         "SK",
	Country:     "Slovakia",
	CountryCode: 421,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "6", Network: "O2"},
		{Code: "1", Network: "Orange"},
		{Code: "7", Network: "Orange"},
		{Code: "5", Network: "Orange"},
		{Code: "3", Network: "Swan / 4ka"},
		{Code: "2", Network: "Telekom"},
		{Code: "50", Network: "Telekom"},
		{Code: "4", Network: "Telekom"},
		{Code: "8", Network: "Uniphone"},
		{Code: "299", Network: "Vonage"},
		{Code: "99", Network: "ZSR"},
	},
}
