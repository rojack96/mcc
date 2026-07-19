package countries

import "github.com/rojack96/mcc/models"

var Colombia = models.Mcc{
	Code:        732,
	Iso:         "CO",
	Country:     "Colombia",
	CountryCode: 57,
	Mnc: []models.Mnc{
		{Code: "101", Network: "Comcel S.A. Occel S.A./Celcaribe"},
		{Code: "102", Network: "Bellsouth Colombia S.A."},
		{Code: "103", Network: "Colombia M"},
		{Code: "111", Network: "TIGO/Colombia Movil"},
		{Code: "123", Network: "Movistar"},
		{Code: "130", Network: "Avantel SAS"},
		{Code: "999", Network: "Fix Line"},
	},
}
