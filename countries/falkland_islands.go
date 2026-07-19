package countries

import "github.com/rojack96/mcc/models"

var FalklandIslands = models.Mcc{
	Code:        750,
	Iso:         "FK",
	Country:     "Falkland Islands (Malvinas)",
	CountryCode: 500,
	Mnc: []models.Mnc{
		{Code: "001", Network: "Cable and Wireless South Atlantic Ltd (Falkland Islands (Malvinas))"},
	},
}
