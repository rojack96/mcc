package countries

import "github.com/rojack96/mcc/models"

var Zimbabwe = models.Mcc{
	Code:        648,
	Iso:         "ZW",
	Country:     "Zimbabwe",
	CountryCode: 263,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Net One"},
		{Code: "03", Network: "Telecel"},
		{Code: "04", Network: "Econet"},
	},
}
