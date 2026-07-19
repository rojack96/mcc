package countries

import "github.com/rojack96/mcc/models"

var Vietnam = models.Mcc{
	Code:        452,
	Iso:         "VN",
	Country:     "Vietnam",
	CountryCode: 84,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Mobifone"},
		{Code: "02", Network: "Vinaphone"},
		{Code: "03", Network: "S-Fone/Telecom"},
		{Code: "04", Network: "Viettel Mobile"},
		{Code: "05", Network: "VietnaMobile"},
		{Code: "06", Network: "Viettel Mobile"},
		{Code: "07", Network: "GTEL Mobile JSC"},
		{Code: "08", Network: "Viettel Mobile"},
	},
}
