package countries

import "github.com/rojack96/mcc/models"

var Somalia = models.Mcc{
	Code:        637,
	Iso:         "SO",
	Country:     "Somalia",
	CountryCode: 252,
	Mnc: []models.Mnc{
		{Code: "00", Network: "Somlink (STG)"},
		{Code: "01", Network: "Telesom"},
		{Code: "04", Network: "Somafone"},
		{Code: "10", Network: "Nationlink"},
		{Code: "19", Network: "HorTel"},
		{Code: "20", Network: "Somnet"},
		{Code: "30", Network: "Golis"},
		{Code: "50", Network: "Hormuud"},
		{Code: "60", Network: "Nationlink"},
		{Code: "71", Network: "Somtel"},
		{Code: "82", Network: "Somtel"},
	},
}
