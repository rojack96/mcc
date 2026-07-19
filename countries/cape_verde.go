package countries

import "github.com/rojack96/mcc/models"

var CapeVerde = models.Mcc{
	Code:        625,
	Iso:         "CV",
	Country:     "Cape Verde",
	CountryCode: 238,
	Mnc: []models.Mnc{
		{Code: "01", Network: "CV Movel"},
		{Code: "02", Network: "T+ Telecom"},
	},
}
