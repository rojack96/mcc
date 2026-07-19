package countries

import "github.com/rojack96/mcc/models"

var Afghanistan = models.Mcc{
	Code:        412,
	Iso:         "AF",
	Country:     "Afghanistan",
	CountryCode: 93,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Afghan Wireless/AWCC"},
		{Code: "20", Network: "Roshan"},
		{Code: "30", Network: "Etisalat"},
		{Code: "40", Network: "Areeba"},
		{Code: "50", Network: "Etisalat"},
		{Code: "80", Network: "Afghan Telecom Corp. (AT)"},
	},
}
