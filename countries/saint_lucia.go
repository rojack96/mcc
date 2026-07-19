package countries

import "github.com/rojack96/mcc/models"

var SaintLucia = models.Mcc{
	Code:        358,
	Iso:         "LC",
	Country:     "Saint Lucia",
	CountryCode: 1758,
	Mnc: []models.Mnc{
		{Code: "110", Network: "Cable & Wireless"},
		{Code: "30", Network: "Cingular Wireless"},
		{Code: "50", Network: "Digicel (St Lucia) Limited"},
	},
}
