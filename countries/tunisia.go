package countries

import "github.com/rojack96/mcc/models"

var Tunisia = models.Mcc{
	Code:        605,
	Iso:         "TN",
	Country:     "Tunisia",
	CountryCode: 216,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Orange"},
		{Code: "02", Network: "Tunisie Telecom"},
		{Code: "03", Network: "Orascom Telecom"},
		{Code: "06", Network: "Tunisie Telecom"},
	},
}
