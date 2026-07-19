package countries

import "github.com/rojack96/mcc/models"

var Slovenia = models.Mcc{
	Code:        293,
	Iso:         "SI",
	Country:     "Slovenia",
	CountryCode: 386,
	Mnc: []models.Mnc{
		{Code: "40", Network: "A1 / Si.mobil"},
		{Code: "20", Network: "Compatel"},
		{Code: "86", Network: "Elektro Gorenjska"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "999", Network: "Fix Line"},
		{Code: "299", Network: "HOT mobil"},
		{Code: "299", Network: "Me2"},
		{Code: "41", Network: "Mobitel"},
		{Code: "299", Network: "Novatel"},
		{Code: "10", Network: "Slovenske zeleznice"},
		{Code: "299", Network: "SoftNET"},
		{Code: "64", Network: "T-2"},
		{Code: "70", Network: "Telemach / Tusmobil"},
	},
}
