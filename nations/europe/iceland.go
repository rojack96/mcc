package europe

import "github.com/rojack96/mcc/models"

var Iceland = models.Mcc{
	Code:        274,
	Iso:         "IS",
	Country:     "Iceland",
	CountryCode: 354,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "11", Network: "NOVA"},
		{Code: "31", Network: "Siminn"},
		{Code: "08", Network: "Siminn"},
		{Code: "01", Network: "Siminn"},
		{Code: "16", Network: "Tismi"},
		{Code: "04", Network: "Viking Wireless"},
		{Code: "12", Network: "Vodafone"},
		{Code: "03", Network: "Vodafone"},
		{Code: "02", Network: "Vodafone"},
	},
}
