package countries

import "github.com/rojack96/mcc/models"

var LaosPDR = models.Mcc{
	Code:        457,
	Iso:         "LA",
	Country:     "Laos P.D.R.",
	CountryCode: 856,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Lao Tel"},
		{Code: "02", Network: "ETL Mobile"},
		{Code: "03", Network: "UNITEL/LAT"},
		{Code: "08", Network: "Tigo/Millicom"},
	},
}
