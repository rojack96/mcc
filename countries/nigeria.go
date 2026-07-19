package countries

import "github.com/rojack96/mcc/models"

var Nigeria = models.Mcc{
	Code:        621,
	Iso:         "NG",
	Country:     "Nigeria",
	CountryCode: 234,
	Mnc: []models.Mnc{
		{Code: "20", Network: "Airtel/ZAIN/Econet"},
		{Code: "25", Network: "Visafone (Nigeria) CDMA"},
		{Code: "30", Network: "MTN"},
		{Code: "40", Network: "M-Tel/Nigeria Telecom. Ltd."},
		{Code: "50", Network: "Glo Mobile"},
		{Code: "60", Network: "ETISALAT"},
		{Code: "99", Network: "Starcomms (Nigeria) CDMA"},
	},
}
