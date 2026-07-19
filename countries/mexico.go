package countries

import "github.com/rojack96/mcc/models"

var Mexico = models.Mcc{
	Code:        334,
	Iso:         "MX",
	Country:     "Mexico",
	CountryCode: 52,
	Mnc: []models.Mnc{
		{Code: "010", Network: "NEXTEL"},
		{Code: "020", Network: "TelCel/America Movil"},
		{Code: "030", Network: "Movistar/Pegaso"},
		{Code: "040", Network: "IUSACell/UneFon"},
		{Code: "050", Network: "IUSACell/UneFon"},
		{Code: "070", Network: "Operadora Unefon SA de CV"},
		{Code: "080", Network: "Operadora Unefon SA de CV"},
		{Code: "090", Network: "NEXTEL"},
	},
}
