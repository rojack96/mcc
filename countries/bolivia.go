package countries

import "github.com/rojack96/mcc/models"

var Bolivia = models.Mcc{
	Code:        736,
	Iso:         "BO",
	Country:     "Bolivia",
	CountryCode: 591,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Nuevatel"},
		{Code: "02", Network: "Entel Pcs"},
		{Code: "03", Network: "TELECEL BOLIVIA"},
	},
}
