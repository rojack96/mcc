package countries

import "github.com/rojack96/mcc/models"

var Cameroon = models.Mcc{
	Code:        624,
	Iso:         "CM",
	Country:     "Cameroon",
	CountryCode: 237,
	Mnc: []models.Mnc{
		{Code: "01", Network: "MTN"},
		{Code: "02", Network: "Orange"},
		{Code: "04", Network: "Nextel"},
	},
}
