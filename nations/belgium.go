package nations

import "github.com/rojack96/mcc/models"

var Belgium = models.Mcc{
	Code:        206,
	Iso:         "BE",
	Country:     "Belgium",
	CountryCode: 32,
	Mnc: []models.Mnc{
		{20, "Base"},
		{28, "BICS"},
		{25, "Dense Air"},
		{23, "Dust Mobile"},
		{33, "Ericsson"},
		{299, "Failed Calls"},
		{299, "FEBO"},
		{999, "Fix Line"},
		{299, "GianCom"},
		{2, "Infrabel"},
		{299, "interactive digital media / IDM"},
		{299, "L-mobi"},
		{99, "Lancelot"},
		{299, "Legos"},
		{6, "Lycamobile"},
		{30, "Mobile Vikings"},
		{10, "Mobistar / Orange"},
		{299, "Nord Connect"},
		{34, "onoff"},
		{299, "PM Factory"},
		{0, "Proximus"},
		{1, "Proximus"},
		{4, "Proximus"},
		{5, "Telenet"},
		{7, "Vectone Mobile"},
		{8, "VOOmobile"},
		{299, "Voxbone / Bandwidth"},
	},
}
