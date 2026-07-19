package countries

import "github.com/rojack96/mcc/models"

var GuineaBissau = models.Mcc{
	Code:        632,
	Iso:         "GW",
	Country:     "Guinea-Bissau",
	CountryCode: 245,
	Mnc: []models.Mnc{
		{Code: "01", Network: "GuineTel"},
		{Code: "02", Network: "SpaceTel"},
		{Code: "03", Network: "Orange"},
	},
}
