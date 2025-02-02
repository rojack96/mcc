package europe

import "github.com/rojack96/mcc/models"

var Latvia = models.Mcc{
	Code:        247,
	Iso:         "LV",
	Country:     "Latvia",
	CountryCode: 371,
	Mnc: []models.Mnc{
		{Code: "05", Network: "Bite"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "01", Network: "LMT"},
		{Code: "10", Network: "LMT"},
		{Code: "299", Network: "Premium Numbers"},
		{Code: "02", Network: "Tele2"},
		{Code: "04", Network: "Tet"},
		{Code: "03", Network: "TRIATEL"},
		{Code: "08", Network: "VENTA Mobile"},
		{Code: "09", Network: "XOmobile"},
	},
}
