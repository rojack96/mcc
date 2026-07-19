package countries

import "github.com/rojack96/mcc/models"

var Iran = models.Mcc{
	Code:        432,
	Iso:         "IR",
	Country:     "Iran",
	CountryCode: 98,
	Mnc: []models.Mnc{
		{Code: "11", Network: "TCI / MCI"},
		{Code: "14", Network: "TKC/KFZO"},
		{Code: "19", Network: "Mobile Telecommunications Company of Esfahan JV-PJS (MTCE)"},
		{Code: "20", Network: "Rightel"},
		{Code: "32", Network: "Taliya"},
		{Code: "35", Network: "MTN/IranCell"},
		{Code: "70", Network: "MTCE (TCI) - Isfahan Celcom"},
		{Code: "999", Network: "Fix Line"},
	},
}
