package countries

import "github.com/rojack96/mcc/models"

var CongoDemRep = models.Mcc{
	Code:        630,
	Iso:         "CD",
	Country:     "Congo, Dem. Rep.",
	CountryCode: 243,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Vodacom"},
		{Code: "02", Network: "ZAIN CelTel"},
		{Code: "05", Network: "SuperCell"},
		{Code: "86", Network: "Orange RDC sarl"},
		{Code: "88", Network: "Yozma Timeturns sprl (YTT)"},
		{Code: "89", Network: "TIGO/Oasis"},
		{Code: "90", Network: "Africell"},
	},
}
