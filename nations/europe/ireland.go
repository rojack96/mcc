package europe

import "github.com/rojack96/mcc/models"

var Ireland = models.Mcc{
	Code:        272,
	Iso:         "IE",
	Country:     "Ireland",
	CountryCode: 353,
	Mnc: []models.Mnc{
		{Code: "05", Network: "3"},
		{Code: "02", Network: "3"},
		{Code: "17", Network: "3"},
		{Code: "299", Network: "BT"},
		{Code: "16", Network: "Carphone Mobile"},
		{Code: "299", Network: "Cubic Telecom"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "13", Network: "Lycamobile"},
		{Code: "07", Network: "Meteor"},
		{Code: "03", Network: "Meteor"},
		{Code: "08", Network: "Meteor"},
		{Code: "299", Network: "Net Feasa"},
		{Code: "299", Network: "OGCIO"},
		{Code: "11", Network: "Tesco Mobile"},
		{Code: "15", Network: "Virgin Media"},
		{Code: "01", Network: "Vodafone"},
	},
}
