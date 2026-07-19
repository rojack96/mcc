package countries

import "github.com/rojack96/mcc/models"

var Benin = models.Mcc{
	Code:        616,
	Iso:         "BJ",
	Country:     "Benin",
	CountryCode: 229,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Libercom"},
		{Code: "02", Network: "Etisalat/MOOV"},
		{Code: "03", Network: "MTN/Spacetel"},
		{Code: "04", Network: "Bell Benin/BBCOM"},
		{Code: "05", Network: "GloMobile"},
	},
}
