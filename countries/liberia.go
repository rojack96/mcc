package countries

import "github.com/rojack96/mcc/models"

var Liberia = models.Mcc{
	Code:        618,
	Iso:         "LR",
	Country:     "Liberia",
	CountryCode: 231,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Lonestar"},
		{Code: "02", Network: "Libercell"},
		{Code: "04", Network: "Comium BVI"},
		{Code: "07", Network: "Orange"},
	},
}
