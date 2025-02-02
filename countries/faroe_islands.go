package countries

import "github.com/rojack96/mcc/models"

var FaroeIslands = models.Mcc{
	Code:        288,
	Iso:         "FO",
	Country:     "Faroe Islands",
	CountryCode: 298,
	Mnc: []models.Mnc{
		{Code: "299", Network: "Failed Calls"},
		{Code: "01", Network: "Faroese Telecom"},
		{Code: "999", Network: "Fix Line"},
		{Code: "02", Network: "Hey / Kall"},
		{Code: "03", Network: "Tosa"},
	},
}
