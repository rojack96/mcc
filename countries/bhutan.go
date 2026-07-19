package countries

import "github.com/rojack96/mcc/models"

var Bhutan = models.Mcc{
	Code:        402,
	Iso:         "BT",
	Country:     "Bhutan",
	CountryCode: 975,
	Mnc: []models.Mnc{
		{Code: "11", Network: "B-Mobile"},
		{Code: "77", Network: "TashiCell"},
	},
}
