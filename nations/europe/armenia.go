package europe

import "github.com/rojack96/mcc/models"

var Armenia = models.Mcc{
	Code:        283,
	Iso:         "AM",
	Country:     "Armenia",
	CountryCode: 374,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Beeline"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "04", Network: "KT"},
		{Code: "10", Network: "Orange"},
		{Code: "05", Network: "Viva-MTS"},
	},
}
