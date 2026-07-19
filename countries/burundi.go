package countries

import "github.com/rojack96/mcc/models"

var Burundi = models.Mcc{
	Code:        642,
	Iso:         "BI",
	Country:     "Burundi",
	CountryCode: 257,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Spacetel / Econet / Leo"},
		{Code: "02", Network: "Africel / Safaris"},
		{Code: "03", Network: "Onatel / Telecel"},
		{Code: "07", Network: "Smart Mobile / LACELL"},
		{Code: "08", Network: "HiTs Telecom"},
		{Code: "82", Network: "Spacetel / Econet / Leo"},
	},
}
