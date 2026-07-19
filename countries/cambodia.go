package countries

import "github.com/rojack96/mcc/models"

var Cambodia = models.Mcc{
	Code:        456,
	Iso:         "KH",
	Country:     "Cambodia",
	CountryCode: 855,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Mobitel/Cam GSM"},
		{Code: "02", Network: "Hello/Smart Mobile"},
		{Code: "03", Network: "QB/Cambodia Adv. Comms."},
		{Code: "05", Network: "Smart Mobile"},
		{Code: "06", Network: "Smart Mobile"},
		{Code: "08", Network: "Metfone"},
		{Code: "09", Network: "Sotelco Ltd (Beeline Cambodia)"},
		{Code: "11", Network: "SEATEL"},
		{Code: "18", Network: "MFone/Camshin"},
	},
}
