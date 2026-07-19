package countries

import "github.com/rojack96/mcc/models"

var Brazil = models.Mcc{
	Code:        724,
	Iso:         "BR",
	Country:     "Brazil",
	CountryCode: 55,
	Mnc: []models.Mnc{
		{Code: "00", Network: "Nextel (Telet)"},
		{Code: "02", Network: "TIM"},
		{Code: "03", Network: "TIM"},
		{Code: "04", Network: "TIM"},
		{Code: "05", Network: "Claro/Albra/America Movil"},
		{Code: "06", Network: "Vivo S.A./Telemig"},
		{Code: "08", Network: "Maxitel MG TIM"},
		{Code: "10", Network: "Vivo S.A./Telemig"},
		{Code: "11", Network: "Vivo S.A./Telemig"},
		{Code: "12", Network: "Americel Claro"},
		{Code: "15", Network: "Sercontel Cel"},
		{Code: "16", Network: "Oi (TNL PCS / Oi)"},
		{Code: "19", Network: "Telemig Cel"},
		{Code: "23", Network: "Vivo S.A./Telemig"},
		{Code: "24", Network: "Oi (TNL PCS / Oi)"},
		{Code: "26", Network: "AmericaNet"},
		{Code: "31", Network: "Oi (TNL PCS / Oi)"},
		{Code: "39", Network: "Nextel (Telet)"},
		{Code: "54", Network: "PORTO SEGURO TELECOMUNICAÇÔES"},
	},
}
