package countries

import "github.com/rojack96/mcc/models"

var PalauRepublicOf = models.Mcc{
	Code:        552,
	Iso:         "PW",
	Country:     "Palau (Republic of)",
	CountryCode: 680,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Palau National Communications Corp. (PNCC) (Palau (Republic of))"},
		{Code: "02", Network: "PECI/PalauTel (Palau (Republic of))"},
	},
}
