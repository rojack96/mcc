package countries

import "github.com/rojack96/mcc/models"

var Lesotho = models.Mcc{
	Code:        651,
	Iso:         "LS",
	Country:     "Lesotho",
	CountryCode: 266,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Vodacom Lesotho"},
		{Code: "02", Network: "Econet/Ezi-cel"},
	},
}
