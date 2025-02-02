package europe

import "github.com/rojack96/mcc/models"

var Luxembourg = models.Mcc{
	Code:        270,
	Iso:         "LU",
	Country:     "Luxembourg",
	CountryCode: 352,
	Mnc: []models.Mnc{
		{Code: "10", Network: "Blue Communications"},
		{Code: "299", Network: "Bouygues Telecom"},
		{Code: "81", Network: "e-LUX Mobile"},
		{Code: "299", Network: "Eltrona"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "05", Network: "Luxembourg Online"},
		{Code: "299", Network: "MTX Connect"},
		{Code: "99", Network: "Orange"},
		{Code: "01", Network: "Post"},
		{Code: "77", Network: "Tango"},
	},
}
