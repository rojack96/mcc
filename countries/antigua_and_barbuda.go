package countries

import "github.com/rojack96/mcc/models"

var AntiguaAndBarbuda = models.Mcc{
	Code:        344,
	Iso:         "AG",
	Country:     "Antigua and Barbuda",
	CountryCode: 1268,
	Mnc: []models.Mnc{
		{Code: "030", Network: "APUA PCS"},
		{Code: "920", Network: "C & W"},
		{Code: "930", Network: "Cing. Wirel./DigiCel"},
	},
}
