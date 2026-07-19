package countries

import "github.com/rojack96/mcc/models"

var Uganda = models.Mcc{
	Code:        641,
	Iso:         "UG",
	Country:     "Uganda",
	CountryCode: 256,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Celtel"},
		{Code: "10", Network: "MTN Ltd."},
		{Code: "11", Network: "Uganda Telecom Ltd."},
		{Code: "14", Network: "Orange"},
		{Code: "18", Network: "Suretelecom Uganda Ltd"},
		{Code: "22", Network: "Warid Telecom"},
		{Code: "30", Network: "K2 Telecom Ltd"},
		{Code: "33", Network: "Smile Communications Uganda Ltd"},
		{Code: "66", Network: "i-Tel Ltd"},
	},
}
