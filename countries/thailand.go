package countries

import "github.com/rojack96/mcc/models"

var Thailand = models.Mcc{
	Code:        520,
	Iso:         "TH",
	Country:     "Thailand",
	CountryCode: 66,
	Mnc: []models.Mnc{
		{Code: "00", Network: "Hutch/CAT CDMA"},
		{Code: "01", Network: "AIS/Advanced Info Service"},
		{Code: "03", Network: "Advanced Wireless Networks/AWN"},
		{Code: "04", Network: "True Move/Orange"},
		{Code: "05", Network: "Total Access (DTAC)"},
		{Code: "15", Network: "ACT Mobile"},
		{Code: "18", Network: "Total Access (DTAC)"},
		{Code: "99", Network: "True Move/Orange"},
		{Code: "999", Network: "Fix Line"},
	},
}
