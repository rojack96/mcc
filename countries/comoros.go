package countries

import "github.com/rojack96/mcc/models"

var Comoros = models.Mcc{
	Code:        654,
	Iso:         "KM",
	Country:     "Comoros",
	CountryCode: 269,
	Mnc: []models.Mnc{
		{Code: "01", Network: "HURI - SNPT"},
		{Code: "02", Network: "TELCO SA"},
	},
}
