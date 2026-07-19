package countries

import "github.com/rojack96/mcc/models"

var CentralAfricanRep = models.Mcc{
	Code:        623,
	Iso:         "CF",
	Country:     "Central African Rep.",
	CountryCode: 236,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Centrafr. Telecom+"},
		{Code: "02", Network: "Telecel Centraf."},
		{Code: "03", Network: "Orange/Celca"},
		{Code: "04", Network: "Nationlink"},
	},
}
