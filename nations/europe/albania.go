package europe

import "github.com/rojack96/mcc/models"

var Albania = models.Mcc{
	Code:        276,
	Iso:         "AL",
	Country:     "Albania",
	CountryCode: 355,
	Mnc: []models.Mnc{
		{Code: "03", Network: "ALBtelecom Mobile / Eagle"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "01", Network: "One / AMC"},
		{Code: "02", Network: "Vodafone"},
	},
}
