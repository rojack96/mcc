package countries

import "github.com/rojack96/mcc/models"

var Belgium = models.Mcc{
	Code:        206,
	Iso:         "BE",
	Country:     "Belgium",
	CountryCode: 32,
	Mnc: []models.Mnc{
		{Code: "20", Network: "Base"},
		{Code: "28", Network: "BICS"},
		{Code: "25", Network: "Dense Air"},
		{Code: "23", Network: "Dust Mobile"},
		{Code: "33", Network: "Ericsson"},
		{Code: "299", Network: "Failed Calls"},
		{Code: "299", Network: "FEBO"},
		{Code: "999", Network: "Fix Line"},
		{Code: "299", Network: "GianCom"},
		{Code: "2", Network: "Infrabel"},
		{Code: "299", Network: "interactive digital media / IDM"},
		{Code: "299", Network: "L-mobi"},
		{Code: "99", Network: "Lancelot"},
		{Code: "299", Network: "Legos"},
		{Code: "6", Network: "Lycamobile"},
		{Code: "30", Network: "Mobile Vikings"},
		{Code: "10", Network: "Mobistar / Orange"},
		{Code: "299", Network: "Nord Connect"},
		{Code: "34", Network: "onoff"},
		{Code: "299", Network: "PM Factory"},
		{Code: "0", Network: "Proximus"},
		{Code: "1", Network: "Proximus"},
		{Code: "4", Network: "Proximus"},
		{Code: "5", Network: "Telenet"},
		{Code: "7", Network: "Vectone Mobile"},
		{Code: "8", Network: "VOOmobile"},
		{Code: "299", Network: "Voxbone / Bandwidth"},
	},
}
