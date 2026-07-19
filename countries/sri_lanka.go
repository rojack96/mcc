package countries

import "github.com/rojack96/mcc/models"

var SriLanka = models.Mcc{
	Code:        413,
	Iso:         "LK",
	Country:     "Sri Lanka",
	CountryCode: 94,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Mobitel Ltd."},
		{Code: "02", Network: "Dialog"},
		{Code: "03", Network: "Etisalat/Tigo"},
		{Code: "05", Network: "Bharti Airtel"},
		{Code: "08", Network: "H3G Hutchison"},
	},
}
