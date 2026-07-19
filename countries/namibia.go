package countries

import "github.com/rojack96/mcc/models"

var Namibia = models.Mcc{
	Code:        649,
	Iso:         "NA",
	Country:     "Namibia",
	CountryCode: 264,
	Mnc: []models.Mnc{
		{Code: "01", Network: "MTC"},
		{Code: "02", Network: "Switch/Nam. Telec."},
		{Code: "03", Network: "Leo / Orascom"},
	},
}
