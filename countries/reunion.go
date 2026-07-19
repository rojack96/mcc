package countries

import "github.com/rojack96/mcc/models"

var Reunion = models.Mcc{
	Code:        647,
	Iso:         "RE",
	Country:     "Reunion",
	CountryCode: 262,
	Mnc: []models.Mnc{
		{Code: "00", Network: "Orange"},
		{Code: "02", Network: "Outremer Telecom"},
		{Code: "10", Network: "SFR"},
	},
}
