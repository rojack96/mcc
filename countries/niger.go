package countries

import "github.com/rojack96/mcc/models"

var Niger = models.Mcc{
	Code:        614,
	Iso:         "NE",
	Country:     "Niger",
	CountryCode: 227,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Sahelcom"},
		{Code: "02", Network: "Zain/CelTel"},
		{Code: "03", Network: "Etisalat/TeleCel"},
		{Code: "04", Network: "Orange Niger SA"},
	},
}
