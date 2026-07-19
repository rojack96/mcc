package countries

import "github.com/rojack96/mcc/models"

var Zambia = models.Mcc{
	Code:        645,
	Iso:         "ZM",
	Country:     "Zambia",
	CountryCode: 260,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Zain/Celtel"},
		{Code: "02", Network: "MTN/Telecel"},
		{Code: "03", Network: "Cell Z/MTS"},
	},
}
