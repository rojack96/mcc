package countries

import "github.com/rojack96/mcc/models"

var IvoryCoast = models.Mcc{
	Code:        612,
	Iso:         "CI",
	Country:     "Ivory Coast",
	CountryCode: 225,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Cora de Comstar"},
		{Code: "02", Network: "Atlantik Tel./Moov"},
		{Code: "03", Network: "Orange"},
		{Code: "04", Network: "Comium"},
		{Code: "05", Network: "MTN"},
		{Code: "06", Network: "OriCell"},
		{Code: "07", Network: "Aircomm SA"},
	},
}
