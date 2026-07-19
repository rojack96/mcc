package countries

import "github.com/rojack96/mcc/models"

var Taiwan = models.Mcc{
	Code:        466,
	Iso:         "TW",
	Country:     "Taiwan",
	CountryCode: 886,
	Mnc: []models.Mnc{
		{Code: "01", Network: "Far EasTone"},
		{Code: "02", Network: "Far EasTone"},
		{Code: "03", Network: "Far EasTone"},
		{Code: "05", Network: "Asia Pacific Telecom Co. Ltd (APT)"},
		{Code: "06", Network: "Far EasTone"},
		{Code: "07", Network: "Far EasTone"},
		{Code: "11", Network: "Chunghwa Telecom LDM"},
		{Code: "88", Network: "KG Telecom /Far EasTone"},
		{Code: "89", Network: "VIBO"},
		{Code: "90", Network: "T-Star/VIBO"},
		{Code: "92", Network: "Chunghwa Telecom LDM"},
		{Code: "93", Network: "Taiwan Cellular"},
		{Code: "97", Network: "Taiwan Cellular"},
		{Code: "99", Network: "Taiwan Cellular"},
	},
}
