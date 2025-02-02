package europe

import "github.com/rojack96/mcc/models"

var Portugal = models.Mcc{
	Code:        268,
	Iso:         "PT",
	Country:     "Portugal",
	CountryCode: 351,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "04", Network: "Lycamobile"},
		{Code: "06", Network: "MEO"},
		{Code: "80", Network: "MEO"},
		{Code: "08", Network: "MEO"},
		{Code: "03", Network: "NOS"},
		{Code: "93", Network: "NOS"},
		{Code: "299", Network: "NOWO"},
		{Code: "299", Network: "Oni"},
		{Code: "91", Network: "Vodafone"},
		{Code: "01", Network: "Vodafone"},
	},
}
