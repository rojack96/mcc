package countries

import "github.com/rojack96/mcc/models"

var SierraLeone = models.Mcc{
	Code:        619,
	Iso:         "SL",
	Country:     "Sierra Leone",
	CountryCode: 232,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Celtel"},
		{Code: "02", Network: "Millicom"},
		{Code: "03", Network: "Africel"},
		{Code: "04", Network: "Comium"},
		{Code: "05", Network: "Lintel"},
		{Code: "07", Network: "Qcell"},
		{Code: "25", Network: "Mobitel"},
	},
}
