package countries

import "github.com/rojack96/mcc/models"

var FrenchGuianaAndGuadeloupe = models.Mcc{
	Code:        340,
	Iso:         "GF/GP",
	Country:     "French Guiana and Guadeloupe",
	CountryCode: 0,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Orange Caribe Mobiles"},
		{Code: "02", Network: "Outremer Telecom"},
		{Code: "03", Network: "TelCell GSM"},
		{Code: "08", Network: "Dauphin Telecom SU (Guadeloupe Telecom) (Guadeloupe (French Department of))"},
		{Code: "11", Network: "TelCell GSM"},
		{Code: "20", Network: "Bouygues/DigiCel"},
	},
}
