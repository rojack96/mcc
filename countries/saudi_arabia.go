package countries

import "github.com/rojack96/mcc/models"

var SaudiArabia = models.Mcc{
	Code:        420,
	Iso:         "SA",
	Country:     "Saudi Arabia",
	CountryCode: 966,
	Mnc: []models.Mnc{
		{Code: "01", Network: "STC/Al Jawal"},
		{Code: "03", Network: "Etihad/Etisalat/Mobily"},
		{Code: "04", Network: "Zain"},
		{Code: "05", Network: "Virgin Mobile"},
		{Code: "06", Network: "Lebara Mobile"},
	},
}
