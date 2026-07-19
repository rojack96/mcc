package countries

import "github.com/rojack96/mcc/models"

var ArgentinaRepublic = models.Mcc{
	Code:        722,
	Iso:         "AR",
	Country:     "Argentina Republic",
	CountryCode: 54,
	Mnc: []models.Mnc{
		{Code: "007", Network: "Movistar/Telefonica"},
		{Code: "020", Network: "Nextel"},
		{Code: "070", Network: "Movistar/Telefonica"},
		{Code: "310", Network: "Claro/ CTI/AMX"},
		{Code: "320", Network: "Compa"},
		{Code: "330", Network: "Compa"},
		{Code: "340", Network: "Telecom Personal S.A."},
		{Code: "341", Network: "Telecom Personal S.A."},
		{Code: "999", Network: "Fix Line"},
	},
}
