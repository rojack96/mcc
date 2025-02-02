package countries

import "github.com/rojack96/mcc/models"

var Liechtenstein = models.Mcc{
	Code:        295,
	Iso:         "LI",
	Country:     "Liechtenstein",
	CountryCode: 423,
	Mnc: []models.Mnc{
		{Code: "02", Network: "7acht"},
		{Code: "06", Network: "CUBIC"},
		{Code: "299", Network: "Datamobile"},
		{Code: "299", Network: "Dimoco"},
		{Code: "09", Network: "EMnify"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "01", Network: "FL GSM"},
		{Code: "05", Network: "FL1"},
		{Code: "299", Network: "SORACOM"},
		{Code: "299", Network: "Telna"},
	},
}
