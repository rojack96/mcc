package countries

import "github.com/rojack96/mcc/models"

var Tanzania = models.Mcc{
	Code:        640,
	Iso:         "TZ",
	Country:     "Tanzania",
	CountryCode: 255,
	Mnc: []models.Mnc{
		{Code: "02", Network: "TIGO/MIC"},
		{Code: "03", Network: "Zantel/Zanzibar Telecom"},
		{Code: "04", Network: "Vodacom Ltd"},
		{Code: "05", Network: "ZAIN/Celtel"},
		{Code: "06", Network: "Dovetel (T) Ltd"},
		{Code: "07", Network: "Tanzania Telecommunications Company Ltd (TTCL)"},
		{Code: "08", Network: "Benson Informatics Ltd"},
		{Code: "09", Network: "ExcellentCom (T) Ltd"},
		{Code: "11", Network: "Smile Communications Tanzania Ltd"},
	},
}
