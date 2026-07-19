package countries

import "github.com/rojack96/mcc/models"

var CuracaoAndNetherlandsAntilles = models.Mcc{
	Code:        362,
	Iso:         "CW/AN",
	Country:     "Curacao and Netherlands Antilles",
	CountryCode: 599,
	Mnc: []models.Mnc{
		{Code: "51", Network: "TELCELL GSM"},
		{Code: "69", Network: "Polycom N.V./ Curacao Telecom d.b.a. Digicel"},
		{Code: "91", Network: "UTS SETEL GSM"},
		{Code: "95", Network: "EOCG Wireless NV"},
		{Code: "951", Network: "UTS Wireless Curacao"},
	},
}
