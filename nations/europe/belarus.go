package europe

import "github.com/rojack96/mcc/models"

var Belarus = models.Mcc{
	Code:        257,
	Iso:         "BY",
	Country:     "Belarus",
	CountryCode: 375,
	Mnc: []models.Mnc{
		{Code: "04", Network: "life:)"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "02", Network: "MTS"},
		{Code: "01", Network: "velcom A1"},
	},
}
