package countries

import "github.com/rojack96/mcc/models"

var Singapore = models.Mcc{
	Code:        525,
	Iso:         "SG",
	Country:     "Singapore",
	CountryCode: 65,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Singtel"},
		{Code: "02", Network: "Singtel"},
		{Code: "03", Network: "MobileOne Ltd"},
		{Code: "05", Network: "Starhub"},
		{Code: "06", Network: "Starhub"},
		{Code: "07", Network: "Singtel"},
		{Code: "10", Network: "Simba"},
		{Code: "999", Network: "Fix Line"},
	},
}
