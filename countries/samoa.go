package countries

import "github.com/rojack96/mcc/models"

var Samoa = models.Mcc{
	Code:        549,
	Iso:         "WS",
	Country:     "Samoa",
	CountryCode: 685,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Telecom Samoa Cellular Ltd."},
		{Code: "27", Network: "Samoatel Mobile"},
		{Code: "999", Network: "Fix Line"},
	},
}
