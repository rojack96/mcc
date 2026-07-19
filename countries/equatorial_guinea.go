package countries

import "github.com/rojack96/mcc/models"

var EquatorialGuinea = models.Mcc{
	Code:        627,
	Iso:         "GQ",
	Country:     "Equatorial Guinea",
	CountryCode: 240,
	Mnc: []models.Mnc{
		{Code: "01", Network: "ORANGE/GETESA"},
		{Code: "03", Network: "Muni"},
	},
}
