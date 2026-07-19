package countries

import "github.com/rojack96/mcc/models"

var Barbados = models.Mcc{
	Code:        342,
	Iso:         "BB",
	Country:     "Barbados",
	CountryCode: 1246,
	Mnc: []models.Mnc{
		{Code: "050", Network: "Digicel"},
		{Code: "600", Network: "C & W BET Ltd."},
		{Code: "750", Network: "Digicel"},
		{Code: "820", Network: "Sunbeach"},
	},
}
