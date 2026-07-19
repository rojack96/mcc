package countries

import "github.com/rojack96/mcc/models"

var Algeria = models.Mcc{
	Code:        603,
	Iso:         "DZ",
	Country:     "Algeria",
	CountryCode: 213,
	Mnc: []models.Mnc{
		{Code: "01", Network: "ATM Mobils"},
		{Code: "02", Network: "Orascom / DJEZZY"},
		{Code: "03", Network: "Oreedo/Wataniya / Nedjma"},
	},
}
