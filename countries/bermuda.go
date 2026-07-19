package countries

import "github.com/rojack96/mcc/models"

var Bermuda = models.Mcc{
	Code:        350,
	Iso:         "BM",
	Country:     "Bermuda",
	CountryCode: 1441,
	Mnc: []models.Mnc{
		{Code: "000", Network: "Bermuda Digital Communications Ltd (BDC)"},
		{Code: "01", Network: "Telecommunications (Bermuda & West Indies) Ltd (Digicel Bermuda)"},
		{Code: "02", Network: "M3 Wireless Ltd"},
		{Code: "10", Network: "Cingular Wireless"},
	},
}
