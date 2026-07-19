package countries

import "github.com/rojack96/mcc/models"

var DominicanRepublic = models.Mcc{
	Code:        370,
	Iso:         "DO",
	Country:     "Dominican Republic",
	CountryCode: 1809,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Orange"},
		{Code: "02", Network: "Claro"},
		{Code: "03", Network: "TRIcom"},
		{Code: "04", Network: "Viva"},
	},
}
