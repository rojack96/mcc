package countries

import "github.com/rojack96/mcc/models"

var Mali = models.Mcc{
	Code:        610,
	Iso:         "ML",
	Country:     "Mali",
	CountryCode: 223,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Malitel"},
		{Code: "02", Network: "Orange/IKATEL"},
		{Code: "03", Network: "ATEL SA"},
	},
}
