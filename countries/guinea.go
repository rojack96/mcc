package countries

import "github.com/rojack96/mcc/models"

var Guinea = models.Mcc{
	Code:        611,
	Iso:         "GN",
	Country:     "Guinea",
	CountryCode: 224,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Orange/Spacetel"},
		{Code: "02", Network: "SotelGui"},
		{Code: "03", Network: "Intercel"},
		{Code: "04", Network: "Areeba"},
		{Code: "05", Network: "Celcom"},
	},
}
